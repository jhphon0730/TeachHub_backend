package main 

import (
	"embed"

	"image_storage_server/router"
)

//go:embed images/*
var images embed.FS

const (
	PORT = "0.0.0.0:8080"
)

func main() {
	if err := router.Runserver(images); err != nil {
		panic(err)
	}
}
