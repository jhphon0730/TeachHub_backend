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

	// ############################ ENrollment ############################
	EnrollmentService = service.NewEnrollmentService()
	EnrollmentHandler = handlers.NewEnrollmentHandler(EnrollmentService)

	// ############################ Dashboard ############################
	DashboardService = service.NewDashboardService()
	DashboardHandler = handlers.NewDashboardHandler(DashboardService)
)

func Runserver() error {
	router := http.NewServeMux()

	// ############################ User ############################
	router.Handle("POST /register", middlewareStack(http.HandlerFunc(UserHandler.RegisterUser)))
	router.Handle("POST /login", middlewareStack(http.HandlerFunc(UserHandler.LoginUser)))
	router.Handle("PATCH /update", middlewareStack(http.HandlerFunc(UserHandler.UpdateUser)))

	// ############################ Course ############################
	router.Handle("POST /course", authMiddlewareStack(http.HandlerFunc(CourseHandler.CreateCourse)))
	router.Handle("DELETE /course/student", authMiddlewareStack(http.HandlerFunc(CourseHandler.RemoveStudentToCourse)))

	// ############################ Enrollment ############################
	router.Handle("GET /enrollment/instructor", authMiddlewareStack(http.HandlerFunc(CourseHandler.GetCourseByInstructorID)))
	router.Handle("GET /enrollment/student", authMiddlewareStack(http.HandlerFunc(EnrollmentHandler.GetCourseByStudentID)))
	router.Handle("POST /enrollment/student", authMiddlewareStack(http.HandlerFunc(EnrollmentHandler.AddStudentEnrollment)))

	// ############################ Dashboard ############################
	router.Handle("GET /dashboard/student/initial", authMiddlewareStack(http.HandlerFunc(DashboardHandler.InitialStudentDashboard)))
	router.Handle("GET /dashboard/instructor/initial", authMiddlewareStack(http.HandlerFunc(DashboardHandler.InitialInstructorDashboard)))

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
