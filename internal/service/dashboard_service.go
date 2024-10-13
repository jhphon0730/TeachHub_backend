package service

import (
	"errors"
	"net/http"

	"image_storage_server/internal/model"
	"image_storage_server/internal/model/dto"
	"image_storage_server/internal/middleware"
)

type DashboardService interface {
	InitialStudentDashboard(r *http.Request) (dto.InitialStudentDashboardDTO, error)
	InitialInstructorDashboard(r *http.Request) (dto.InitialInstructorDashboardDTO, error)
}

type dashboardService struct { }

func NewDashboardService() DashboardService {
	return &dashboardService{}
}

func (c *dashboardService) InitialStudentDashboard(r *http.Request) (dto.InitialStudentDashboardDTO, error) { 
	var initialDTO dto.InitialStudentDashboardDTO
	var err error

	user, ok := r.Context().Value(middleware.UserContextKey).(*model.User)
	if !ok || user == nil {
		return initialDTO, errors.New("User not found")
	}
	if user.Role != "student" {
		return initialDTO, errors.New("User is not student")
	}

	stdCount, err := model.GetAllStudentsCount()
	if err != nil {
		return initialDTO, errors.New("Failed to get all students count")
	}
	insCount, err := model.GetAllInstructorsCount()
	if err != nil {
		return initialDTO, errors.New("Failed to get all instructors count")
	}
	courseCount, err := model.GetAllCoursesCount()
	if err != nil {
		return initialDTO, errors.New("Failed to get all courses count")
	}
	myCourseCount, err := model.GetAllMyCoursesCountByStudentID(user.ID)

	initialDTO.TotalStudentCount = stdCount
	initialDTO.TotalInstructorCount = insCount
	initialDTO.TotalCourseCount = courseCount
	initialDTO.MyCourseCount = myCourseCount

	return initialDTO, err
}

func (c *dashboardService) InitialInstructorDashboard(r *http.Request) (dto.InitialInstructorDashboardDTO, error) { 
	var initialDTO dto.InitialInstructorDashboardDTO
	var err error

	user, ok := r.Context().Value(middleware.UserContextKey).(*model.User)
	if !ok || user == nil {
		return initialDTO, errors.New("User not found")
	}
	if user.Role != "instructor" {
		return initialDTO, errors.New("User is not instructor")
	}

	stdCount, err := model.GetAllStudentsCount()
	if err != nil {
		return initialDTO, errors.New("Failed to get all students count")
	}
	insCount, err := model.GetAllInstructorsCount()
	if err != nil {
		return initialDTO, errors.New("Failed to get all instructors count")
	}
	courseCount, err := model.GetAllCoursesCount()
	if err != nil {
		return initialDTO, errors.New("Failed to get all courses count")
	}
	myCourseCount, err := model.GetAllMyCoursesCountByInstructorID(user.ID)
	if err != nil {
		return initialDTO, errors.New("Failed to get all my courses count")
	}

	initialDTO.TotalStudentCount = stdCount
	initialDTO.TotalInstructorCount = insCount
	initialDTO.TotalCourseCount = courseCount
	initialDTO.MyCourseCount = myCourseCount

	return initialDTO, err
}
