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
	// Create Handlers
	r.newImageHandler()

	// Create Middleware 
	middlewareStack := middleware.CreateMiddlewareStack(
		middleware.CORS,
		middleware.Logger,
	)


	server := &http.Server{
		Addr:    PORT,
		Handler: middlewareStack(r.router),
	}

	log.Println("Server is running on port", PORT)
	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}

	http.ListenAndServe(PORT, r.router)
}

func (r *Router) newImageHandler() {
	r.HandleFunc("POST /upload", ImageController.SaveImage)
	// get query string 
		// ex) http://localhost:8080/read?imageName=sample.jpg
	r.HandleFunc("GET /read", ImageController.ReadImage)
}

func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.router.HandleFunc(pattern, handler)
}

