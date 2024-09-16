package service

import (
	"net/http"
	"io"
	"os"
	"time"
	"errors"
)

type ImageService interface {
	ReadImage(r *http.Request) ([]byte, time.Time, error)
	SaveImage(r *http.Request) error
}

type imageService struct{}

// NewImageService는 ImageService를 반환합니다.
func NewImageService() ImageService {
	return &imageService{}
}

// ReadImage는 이미지 파일을 읽고 바이트 데이터와 수정 시간을 반환합니다.
func (s *imageService) ReadImage(r *http.Request) ([]byte, time.Time, error) {
	imageName := r.URL.Query().Get("imageName")
	if len(imageName) == 0 {
		return nil, time.Time{}, errors.New("imageName is required")
	}

	file, err := os.Open("images/" + imageName)
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

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		return err
	}
	defer file.Close()

	filename := fileHeader.Filename
	if _, err := os.Stat("images/" + filename); err == nil {
		return nil // 파일이 이미 존재함
	}

	dst, err := os.Create("images/" + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}
