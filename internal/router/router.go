package router

import (
	"log"
	"net/http"
	"os"
	"embed"

	"image_storage_server/internal/router/middleware"
	"image_storage_server/internal/handler"
)

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080" // 기본 포트 설정
	}
	return ":" + port
}

func Runserver(imagesEmbed embed.FS) error {
	router := http.NewServeMux()

	ImageHandler := handler.NewImageHandler(imagesEmbed)

	// 메서드를 직접 라우팅에 포함
	router.HandleFunc("POST /upload", ImageHandler.SaveImage)
	router.HandleFunc("GET /read", ImageHandler.ReadImage)

	// 미들웨어 스택 생성
	middlewareStack := middleware.ChainMiddleware(
		middleware.CORS,
		middleware.Logger,
	)

	// 서버 설정
	server := &http.Server{
		Addr:    getPort(),
		Handler: middlewareStack(router),
	}

	log.Println("Server is running on port", getPort())
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}

	return nil
}
