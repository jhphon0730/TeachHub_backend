package main

import (
	"embed"
	"log"

	"image_storage_server/internal/router"
	"image_storage_server/config"
)

//go:embed images/*
var images embed.FS

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load the environment variables: %v", err)
	}


	if err := router.Runserver(images); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
