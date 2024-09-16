package main

import (
	"log"

	"image_storage_server/internal/router"
	"image_storage_server/config"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load the environment variables: %v", err)
	}


	if err := router.Runserver(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
