package service

import (
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/model"
)

type CoursesService interface {
	CreateCourses(r *http.Request) (error)
}

type coursesService struct { }

func NewCoursesService() CoursesService {
	return &coursesService{}
}

func (c *coursesService) CreateCourses(r *http.Request) error {
	var courses model.Courses
	var err error

	if err = utils.ParseJSON(r, &courses); err != nil {
		return err
	}

	// Valid Input 
	if err = utils.CheckValidCreateCoursesInput(&courses); err != nil {
		return err
	}

	// Check Input Length 
	_, err = model.InsertCourse(&courses)
	if err != nil {
		return err
	}

	return nil
}
