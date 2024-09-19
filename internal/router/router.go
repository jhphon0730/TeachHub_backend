package router

import (
	"log"
	"net/http"


	"image_storage_server/internal/context"
	"image_storage_server/internal/router/middleware"
	"image_storage_server/internal/handlers"
	"image_storage_server/internal/service"
	"image_storage_server/config"
)

var (
	ImageService = service.NewImageService(config.GetImageStorageDir())
	ImageHandler = handlers.NewImageHandler(ImageService)

	UserService = service.NewUserService()
	UserHandler = handlers.NewUserHandler(UserService)
)

func Runserver() error {
	router := http.NewServeMux()

	router.Handle("POST /upload", context.AppHandler{HandleFunc: ImageHandler.SaveImage})
	router.Handle("GET /read", context.AppHandler{HandleFunc: ImageHandler.ReadImage})

	router.Handle("POST /register", context.AppHandler{HandleFunc: UserHandler.RegisterUser})
	router.Handle("POST /login", context.AppHandler{HandleFunc: UserHandler.LoginUser})
	router.Handle("GET /find", context.AppHandler{HandleFunc: UserHandler.FindUser})

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
