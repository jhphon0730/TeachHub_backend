package controller

import (
	"log"
	"net/http"
	"os"
	"io"
	"image_storage_server/json"
)

// struct & interface
type ImageController struct{}

type ImageControllerInterface interface {
	ReadImage(w http.ResponseWriter, r *http.Request)
	SaveImage(w http.ResponseWriter, r *http.Request)
}

// 이렇게 생성한 함수는 interface를 구현하게 된다.
// 이렇게 하면 다른 struct에서도 이 interface를 구현하게 되면
// 이 함수를 사용할 수 있게 된다.
func NewImageController() ImageControllerInterface {
	return &ImageController{}
}

func (ic *ImageController) ReadImage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
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

	// Check the File is Already Exists
	//	if the file is already exists, return file path 
	if _, err := os.Stat(handler.Filename); err == nil {
		log.Println("File is already exists")
		json.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"url": "http://localhost:8080/images/" + handler.Filename,
		})
		return
	}

	// Create a new file in the server
	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file to the server
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"url": "http://localhost:8080/images/" + handler.Filename,
	})
}
