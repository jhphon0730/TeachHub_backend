package handlers

import (
	"net/http"
	"bytes"

	"image_storage_server/internal/service"
	"image_storage_server/pkg/utils"
)

type ImageHandler struct {
	service service.ImageService
}

func NewImageHandler(svc service.ImageService) *ImageHandler {
	return &ImageHandler{service: svc}
}

func (h *ImageHandler) ReadImage(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("imageName")

	if len(imageName) == 0 {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "No image name")
		return
	}

	data, modTime, err := h.service.ReadImage(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "No image")
		return
	}

	http.ServeContent(w, r, imageName, modTime, bytes.NewReader(data))
}

func (h *ImageHandler) SaveImage(w http.ResponseWriter, r *http.Request) {
	err := h.service.SaveImage(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to save image")
		return
	}

	utils.WriteSuccessResponse(w, http.StatusCreated, "Image Created", nil)
}
