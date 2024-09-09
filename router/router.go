package router 

import (
	"log"
	"os"
	"net/http"

	"image_storage_server/middleware"
	"image_storage_server/controller"
)

var (
	ImageController = controller.NewImageController()
)

type Router struct {
	router *http.ServeMux
}

func NewRouter() *Router {
	router := http.NewServeMux()
	Router := &Router{router}

	return Router
} 

func (r *Router) Runserver(PORT string) {
	r.HandleFunc("/upload", ImageController.SaveImage)
	r.HandleFunc("/", ImageController.ReadImage)

	server := &http.Server{
		Addr:    PORT,
		Handler: middleware.CORS(r.router),
	}
	// 이후에 CORS와 같은 미들웨어를 2개 이상 사용할 때는 순서에 주의해야 함
	// CORS(router) -> Logger(CORS(router)) 이런식으로 사용해야 함

	log.Println("Server is running on port", PORT)
	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}

	http.ListenAndServe(PORT, r.router)
}

func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.router.HandleFunc(pattern, handler)
}

