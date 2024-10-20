package service

import (
	"errors"
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/model"
	"image_storage_server/internal/model/dto"
	"image_storage_server/internal/middleware"
)

type CourseService interface {
	CreateCourse(r *http.Request) (error)
	GetCourseByInstructorID(r *http.Request) ([]dto.FindCourseByInstructorIDDTO, error)
	RemoveStudentToCourse(r *http.Request) (error)
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

// ####################### Courses ####################### //
/* 강사의 ID로 강의 조회 */
func (c *courseService) GetCourseByInstructorID(r *http.Request) ([]dto.FindCourseByInstructorIDDTO, error) {
	_, ok := r.Context().Value(middleware.UserContextKey).(*model.User)
	if !ok {
		return nil, errors.New("User not found")
	}
	
	instructor_id := r.URL.Query().Get("instructor_id")
	if len(instructor_id) == 0 {
		return nil, errors.New("instructor_id is empty")
	}

	// id type is int64
	instructorID, err := utils.ParseInt64(instructor_id)
	user, err := model.FindUserByID(instructorID)
	if err != nil || user == nil {
		return nil, errors.New("User not found")
	}

	// Check if the user is an instructor
	if user.Role != "instructor" {
		return nil, errors.New("User is not an instructor")
	}

	courses, err := model.FindCourseByInstructorID(user.ID)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

/* 강사가 학생의 ID로 수강 내역 삭제 */
func (c *courseService) RemoveStudentToCourse(r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserContextKey).(*model.User)
	if !ok || user == nil {
		return errors.New("User not found")
	}

	// Check roll
	if user.Role != "instructor" {
		return errors.New("User is not an instructor")
	}

	var removeStudentDTO dto.RemoveStudentDTO
	var err error
	if err = utils.ParseJSON(r, &removeStudentDTO); err != nil {
		return err
	}

	if removeStudentDTO.Student_Username == "" {
		return errors.New("Student_Username is empty")
	}

	student, err := model.FindUserByUserName(removeStudentDTO.Student_Username)
	if err != nil {
		return errors.New("Cannot find student")
	}

	if student.Role != "student" {
		return errors.New("User is not a student")
	}

	err = model.DeleteEnrollmentByStudentIDAndCourseID(student.ID, removeStudentDTO.Course_id)
	if err != nil {
		return errors.New("Cannot remove student to course")
	}

	return nil
}

