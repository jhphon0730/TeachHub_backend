package main 

import (
	"image_storage_server/router"
)

const (
	PORT = "0.0.0.0:8080"
)

func main() {
	router := router.NewRouter()

	router.Runserver(PORT)
}
