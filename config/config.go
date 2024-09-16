package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(filePaths ...string) error {
	// 환경 변수 파일 경로가 주어지지 않으면 기본 `.env` 파일 로드
	if len(filePaths) == 0 {
		filePaths = append(filePaths, ".env")
	}

	// 환경 변수 로드 시도
	if err := godotenv.Load(filePaths...); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}

	log.Println(".env file loaded successfully")
	return nil
}
