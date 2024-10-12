package handlers

import (
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/service"
)

type DashboardHandler struct {
	service service.DashboardService
}

func NewDashboardHandler(service service.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}

func (d *DashboardHandler) InitialStudentDashboard(w http.ResponseWriter, r *http.Request) {
	initialDashboard, err := d.service.InitialStudentDashboard(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "Initial student dashboard retrieved successfully", initialDashboard)
}

func (d *DashboardHandler) InitialInstructorDashboard(w http.ResponseWriter, r *http.Request) {
	initialDashboard, err := d.service.InitialInstructorDashboard(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "Initial instructor dashboard retrieved successfully", initialDashboard)
}
