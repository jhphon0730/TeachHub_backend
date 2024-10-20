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

func (c *CourseHandler) GetCourseByInstructorID(w http.ResponseWriter, r *http.Request) {
	courses, err := c.service.GetCourseByInstructorID(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "Courses retrieved successfully", courses)
}

func (c *CourseHandler) RemoveStudentToCourse(w http.ResponseWriter, r *http.Request) {
	err := c.service.RemoveStudentToCourse(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "Student removed successfully", nil)
}
