package service

import (
	"errors"
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/model"
	"image_storage_server/internal/middleware"
)

type CourseService interface {
	CreateCourse(r *http.Request) (error)
}

type courseService struct { }

func NewCourseService() CourseService {
	return &courseService{}
}

func (c *courseService) CreateCourse(r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserContextKey).(*model.User)
	if !ok || user == nil {
		return errors.New("User not found")
	}

	var courses model.Courses
	var err error

	if err = utils.ParseJSON(r, &courses); err != nil {
		return err
	}
	courses.Instructor_id = user.ID

	// Valid Input 
	if err = utils.CheckValidCreateCourseInput(&courses); err != nil {
		return err
	}

	// Check Input Length 
	_, err = model.InsertCourse(&courses)
	if err != nil {
		return err
	}

	return nil
}
