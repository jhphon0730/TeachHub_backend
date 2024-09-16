package router

import (
	"log"
	"net/http"

	"image_storage_server/internal/router/middleware"
	"image_storage_server/internal/handlers"
	"image_storage_server/internal/service"
	"image_storage_server/config"
)

func Runserver() error {
	router := http.NewServeMux()

	ImageService := service.NewImageService(config.GetImageStorageDir())
	ImageHandler := handlers.NewImageHandler(ImageService)
	router.HandleFunc("POST /upload", ImageHandler.SaveImage)
	router.HandleFunc("GET /read", ImageHandler.ReadImage)

	// 미들웨어 스택 생성
	middlewareStack := middleware.ChainMiddleware(
		middleware.CORS,
		middleware.Logger,
	)

	// 서버 설정
	server := &http.Server{
		Addr:    config.GetPort(),
		Handler: middlewareStack(router),
	}

	log.Println("Server is running on port", config.GetPort())
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}

	return nil
}
