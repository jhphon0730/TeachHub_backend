package controller

import (
	"net/http"
	"embed"
	"io"
	"bytes"

	"image_storage_server/json"
	"image_storage_server/service"
)

// struct & interface
type ImageController struct{
	imagesEmbed embed.FS

	imageService service.ImageServiceInterface
}

type ImageControllerInterface interface {
	ReadImage(w http.ResponseWriter, r *http.Request)
	SaveImage(w http.ResponseWriter, r *http.Request)
}

func NewImageController(imagesEmbed embed.FS) ImageControllerInterface {
	imageService := service.NewImageService()

	return &ImageController{
		imagesEmbed: imagesEmbed,

		imageService: imageService,
	}
}

func (ic *ImageController) ReadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")

	image_name := r.URL.Query().Get("imageName")

	if len(image_name) == 0 {
		json.ResponseError(w, http.StatusBadRequest, "No image name")
		return
	}

	file, err := ic.imagesEmbed.Open("images/test.png")
	if err != nil {
		json.ResponseError(w, http.StatusBadRequest, "No image")
		return
	}
	defer file.Close()


	// Get the file info
	fileStat, err := file.Stat()
	if err != nil {
		json.ResponseError(w, http.StatusInternalServerError, "No image")
		return
	}

	// Read ( IO Data ) 
	fileData, err := io.ReadAll(file)
	if err != nil {
		json.ResponseError(w, http.StatusInternalServerError, "No image")
		return
	}

	reader := bytes.NewReader(fileData)

	http.ServeContent(w, r, image_name, fileStat.ModTime(), reader)
}

func (ic *ImageController) SaveImage(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		json.ResponseError(w, http.StatusBadRequest, "Image size too large")
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("image")
	if err != nil {
		json.ResponseError(w, http.StatusBadRequest, "Error Retrieving the File")
		return
	}
	defer file.Close()

	// service here
	result, err := ic.imageService.SaveImage(handler.Filename, file)
	if err != nil {
		json.ResponseError(w, http.StatusInternalServerError, "Error Saving the File")
		return
	}

	json.ResponseJSON(w, http.StatusOK, result)
	return
}
