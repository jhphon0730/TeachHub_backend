package router 

import (
	"log"
	"os"
	"net/http"
	"embed"

	"image_storage_server/webServer/router/middleware"
	"image_storage_server/webServer/handler"
)

var (
	PORT = ":8080" // TODO: To env file
)

func Runserver(imagesEmbed embed.FS) error {
	router := http.NewServeMux() 

	ImageHandler := handler.NewImageHandler(imagesEmbed)
	router.HandleFunc("POST /upload", ImageHandler.SaveImage)
	router.HandleFunc("GET /read", ImageHandler.ReadImage)

	// Create Middleware 
	middlewareStack := middleware.CreateMiddlewareStack(
		middleware.CORS,
		middleware.Logger,
	)


	server := &http.Server{
		Addr:    PORT,
		Handler: middlewareStack(router),
	}

	log.Println("Server is running on port", PORT)
	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}

	return http.ListenAndServe(PORT, router)
}

