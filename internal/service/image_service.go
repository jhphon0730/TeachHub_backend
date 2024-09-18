package service

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"image_storage_server/pkg/fs"
)

type ImageService interface {
	ReadImage(r *http.Request) ([]byte, time.Time, error)
	SaveImage(r *http.Request) error
}

type imageService struct {
	storageDir string
}

func NewImageService(storageDir string) ImageService {
	return &imageService{storageDir: storageDir}
}

func (s *imageService) ensureStorageDirExists() error {
	if !fs.DirectoryExists(s.storageDir) {
		if err := os.MkdirAll(s.storageDir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (s *imageService) ReadImage(r *http.Request) ([]byte, time.Time, error) {
	imageName := r.URL.Query().Get("imageName")
	if len(imageName) == 0 {
		return nil, time.Time{}, errors.New("imageName is required")
	}

	filePath := filepath.Join(s.storageDir, imageName)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, time.Time{}, err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return nil, time.Time{}, err
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, time.Time{}, err
	}

	return fileData, fileStat.ModTime(), nil
}

func (s *imageService) SaveImage(r *http.Request) error {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return err
	}

	if err := s.ensureStorageDirExists(); err != nil {
		return err
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		return err
	}
	defer file.Close()

	filename := fileHeader.Filename
	if strings.Contains(filename, "..") {
		return errors.New("invalid file name")
	}

	filePath := filepath.Join(s.storageDir, filename)
	if fs.FileExists(filePath) {
		return errors.New("file already exists")
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}
