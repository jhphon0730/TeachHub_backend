package router

import (
	"log"
	"net/http"
	"time"

	"image_storage_server/config"
	"image_storage_server/internal/middleware"
	"image_storage_server/internal/handlers"
	"image_storage_server/internal/service"
)

var (
	// Default Middleware Stacks
	middlewareStack = middleware.ChainMiddleware(
		middleware.CORS,
		middleware.Logger,
	)
	authMiddlewareStack = middleware.ChainMiddleware(
		middleware.CORS,
		middleware.Logger,
		middleware.Auth,
	)

	// ############################ User ############################
	UserService = service.NewUserService()
	UserHandler = handlers.NewUserHandler(UserService)

	// ############################ Course ############################
	CourseService = service.NewCourseService()
	CourseHandler = handlers.NewCourseHandler(CourseService)
)

func Runserver() error {
	router := http.NewServeMux()

	// ############################ User ############################
	router.Handle("/register", middlewareStack(http.HandlerFunc(UserHandler.RegisterUser)))
	router.Handle("/login", middlewareStack(http.HandlerFunc(UserHandler.LoginUser)))
	router.Handle("/update", middlewareStack(http.HandlerFunc(UserHandler.UpdateUser)))

	// ############################ Course ############################
	router.Handle("/course", authMiddlewareStack(http.HandlerFunc(CourseHandler.CreateCourse)))

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
