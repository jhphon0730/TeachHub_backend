package handlers

import (
	"net/http"
	"bytes"

	"image_storage_server/internal/context"
	"image_storage_server/internal/service"
	"image_storage_server/pkg/utils"
)

type ImageHandler struct {
	service service.ImageService
}

func NewImageHandler(svc service.ImageService) *ImageHandler {
	return &ImageHandler{service: svc}
}

func (h *ImageHandler) ReadImage(c *context.Context) {
	imageName := c.Request.URL.Query().Get("imageName")

	if len(imageName) == 0 {
		utils.WriteErrorResponse(c.Writer, http.StatusBadRequest, "No image name")
		return
	}

	data, modTime, err := h.service.ReadImage(c.Request)
	if err != nil {
		utils.WriteErrorResponse(c.Writer, http.StatusBadRequest, "No image")
		return
	}

	http.ServeContent(c.Writer, c.Request, imageName, modTime, bytes.NewReader(data))
}

func (h *ImageHandler) SaveImage(c *context.Context) {
	err := h.service.SaveImage(c.Request)
	if err != nil {
		utils.WriteErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to save image")
		return
	}

	utils.WriteSuccessResponse(c.Writer, http.StatusCreated, "Image Created", nil)
}
