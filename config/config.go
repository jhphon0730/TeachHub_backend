package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(filePaths ...string) error {
	if len(filePaths) == 0 {
		filePaths = append(filePaths, ".env")
	}

	if err := godotenv.Load(filePaths...); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}

	log.Println(".env file loaded successfully")
	return nil
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func GetImageStorageDir() string {
	storageDir := os.Getenv("IMAGE_STORAGE_DIR")
	if storageDir == "" {
		storageDir = "images/"
	}
	return storageDir
}
