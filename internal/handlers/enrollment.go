package handlers

import (
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/service"
)

type EnrollmentHandler struct {
	service service.EnrollmentService
}

func NewEnrollmentHandler(service service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{
		service: service,
	}
}

func (c *EnrollmentHandler) AddStudentEnrollment(w http.ResponseWriter, r *http.Request) {
	err := c.service.AddStudentEnrollment(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusCreated, "Enrollment added successfully", nil)
}

func (c *EnrollmentHandler) GetCourseByStudentID(w http.ResponseWriter, r *http.Request) {
	courses, err := c.service.GetCourseByStudentID(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "Courses retrieved successfully", courses)
}
