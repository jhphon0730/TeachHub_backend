package router

import (
	"log"
	"net/http"
	"time"

	"image_storage_server/internal/middleware"
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

	router.HandleFunc("POST /upload", ImageHandler.SaveImage)
	router.HandleFunc("GET /read", ImageHandler.ReadImage)

	router.HandleFunc("POST /register", UserHandler.RegisterUser)
	router.HandleFunc("POST /login", UserHandler.LoginUser)
	router.HandleFunc("PUT /user", UserHandler.UpdateUser)


	// 미들웨어 스택 생성
	middlewareStack := middleware.ChainMiddleware(
		middleware.CORS,
		middleware.Logger,
	)

	// 서버 설정
	server := &http.Server{
		Addr:    config.GetPort(),
		Handler: middlewareStack(router),
		ReadTimeout: time.Second * 3, // 읽기 시간 초과
		WriteTimeout: time.Second * 3, // 쓰기 시간 초과
		IdleTimeout: time.Second * 60, // 핸들러가 리턴되기까지 대기하는 시간
	}

	log.Println("Server is running on port", config.GetPort())
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}

	return nil
}
