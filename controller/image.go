package controller

import (
	"net/http"
	"image_storage_server/json"
	"image_storage_server/service"
)

// struct & interface
type ImageController struct{
	imageService service.ImageServiceInterface
}

type ImageControllerInterface interface {
	ReadImage(w http.ResponseWriter, r *http.Request)
	SaveImage(w http.ResponseWriter, r *http.Request)
}

func NewImageController() ImageControllerInterface {
	imageService := service.NewImageService()

	return &ImageController{
		imageService: imageService,
	}
}

func (ic *ImageController) ReadImage(w http.ResponseWriter, r *http.Request) {
	image_name := r.URL.Path[1:]
	if len(image_name) == 0 {
		http.Error(w, "No image name", http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, image_name)
}

func (ic *ImageController) SaveImage(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// service here
	result, err := ic.imageService.SaveImage(handler.Filename, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.ResponseJSON(w, http.StatusOK, result)
	return
}
