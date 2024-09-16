package handlers

import (
	"net/http"
	"embed"
	"io"
	"log"
	"os"
	"bytes"
	"image_storage_server/pkg/utils"
)

// struct & interface
type ImageHandler struct{
	imagesEmbed embed.FS
}

func NewImageHandler(imagesEmbed embed.FS) *ImageHandler {
	return &ImageHandler{
		imagesEmbed: imagesEmbed,
	}
}

func (ic *ImageHandler) ReadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")

	image_name := r.URL.Query().Get("imageName")

	if len(image_name) == 0 {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "No image name")
		return
	}

	file, err := ic.imagesEmbed.Open("images/" + image_name)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "No image")
		return
	}
	defer file.Close()


	// Get the file info
	fileStat, err := file.Stat()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "No image")
		return
	}

	// Read ( IO Data ) 
	fileData, err := io.ReadAll(file)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "No image")
		return
	}

	reader := bytes.NewReader(fileData)

	http.ServeContent(w, r, image_name, fileStat.ModTime(), reader)
}

func (ic *ImageHandler) SaveImage(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Image size too large")
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("image")
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Error Retrieving the File")
		return
	}
	defer file.Close()

	filename := handler.Filename

	// Check the File is Already Exists
	//	if the file is already exists, return file path 
	if _, err := os.Stat(filename); err == nil {
		log.Println("File is already exists")
		response := map[string]interface{}{ "url": "http://localhost:8080/read?imageName=" + filename }
		utils.WriteJSONResponse(w, http.StatusOK, response)
		return
	}

	// Create a new file in the server ( images/filename )
	dst, err := os.Create("images/" + filename)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Cannot create file")
		return
	}
	defer dst.Close()

	// Copy the file to the server
	if _, err := io.Copy(dst, file); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Cannot copy file")
		return
	}

	response := map[string]interface{}{ "url": "http://localhost:8080/read?imageName=" + filename }
	utils.WriteJSONResponse(w, http.StatusOK, response)
	return
}
