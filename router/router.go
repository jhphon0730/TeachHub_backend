package router 

import (
	"log"
	"os"
	"net/http"
	"embed"

	"image_storage_server/middleware"
	"image_storage_server/controller"
)

var (
	PORT = ":8080" // TODO: To env file

	ImageController controller.ImageControllerInterface
)

func Runserver(imagesEmbed embed.FS) error {
	router := http.NewServeMux() 

	ImageController = controller.NewImageController(imagesEmbed)
	router.HandleFunc("POST /image/upload", ImageController.SaveImage)
	router.HandleFunc("GET /image/read", ImageController.ReadImage)



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

