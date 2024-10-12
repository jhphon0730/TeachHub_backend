package utils

import (
	"errors"

	"image_storage_server/internal/model"
)

// check Valid Input [Create]
func CheckValidCreateCourseInput(course *model.Courses) error {
	if course.Instructor_id == 0 {
		return errors.New("Instructor_id is required")
	}
	if course.Title == "" {
		return errors.New("Title is required")
	}
	if course.Description == "" {
		return errors.New("Description is required")
	}
	return nil
}
