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
	ImageController controller.ImageControllerInterface
)

type Router struct {
	router *http.ServeMux
}

func NewRouter(imagesEmbed embed.FS) *Router {
	router := http.NewServeMux()
	Router := &Router{router}

	// Create Controller
	ImageController = controller.NewImageController(imagesEmbed)

	return Router
} 

func (r *Router) Runserver(PORT string) {
	r.newImageHandler()

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

func (r *Router) newImageHandler() {
	r.HandleFunc("/upload", ImageController.SaveImage)
	// get query string 
		// ex) http://localhost:8080/read?imageName=sample.jpg
	r.HandleFunc("/read", ImageController.ReadImage)
}

func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.router.HandleFunc(pattern, handler)
}

