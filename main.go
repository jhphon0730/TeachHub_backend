package main

import (
	"log"

	"image_storage_server/config"
	"image_storage_server/internal/model"
	"image_storage_server/internal/router"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load the environment variables: %v", err)
	}


	if err := router.Runserver(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	if err := model.Connect(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return
}
