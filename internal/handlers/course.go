package handlers

import (
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/service"
)

type CourseHandler struct {
	service service.CourseService
}

func NewCourseHandler(service service.CourseService) *CourseHandler {
	return &CourseHandler{
		service: service,
	}
}

func (c *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	err := c.service.CreateCourse(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusCreated, "Course created successfully", nil)
}
