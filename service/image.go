package service

import (
	"log"
	"os"
	"io"
)
 
type ImageService struct{}

type ImageServiceInterface interface {
	SaveImage(filename string, file io.Reader) (map[string]interface{}, error)
}

func NewImageService() ImageServiceInterface {
	return &ImageService{}
}

func (is *ImageService) SaveImage(filename string, file io.Reader) (map[string]interface{}, error) {
	// Check the File is Already Exists
	//	if the file is already exists, return file path 
	if _, err := os.Stat(filename); err == nil {
		log.Println("File is already exists")
		return map[string]interface{}{ "url": "http://localhost:8080/images/" + filename }, nil
	}

	// Create a new file in the server ( images/filename )
	dst, err := os.Create("images/" + filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy the file to the server
	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}

	return map[string]interface{}{ "url": "http://localhost:8080/images/" + filename }, nil
}
