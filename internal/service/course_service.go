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
	GetCourseByInstructorID(r *http.Request) ([]model.Courses, error)

	AddStudentEnrollment(r *http.Request) (error)
	GetCourseByStudentID(r *http.Request) ([]model.Courses, error)
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
func (c *courseService) GetCourseByInstructorID(r *http.Request) ([]model.Courses, error) {
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

// ####################### Enrollments ####################### //
/* 강사가 학생의 ID로 수강 내역 추가 */
func (c *courseService) AddStudentEnrollment(r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserContextKey).(*model.User)
	if !ok || user == nil {
		return errors.New("User not found")
	}

	var addStudentDTO dto.AddStudentDTO
	var err error
	if err = utils.ParseJSON(r, &addStudentDTO); err != nil {
		return err
	}

	// Check User is instructor
	if user.Role != "instructor" {
		return errors.New("User is not an instructor")
	}

	if addStudentDTO.Course_id == 0 || addStudentDTO.Student_id == 0 {
		return errors.New("Course_id or Student_id is empty")
	}

	_, err = model.InsertStudentEnrollment(addStudentDTO.Course_id, addStudentDTO.Student_id)
	if err != nil {
		return errors.New("Cannot add student to course")
	}

	return nil
}

/* 학생 ID로 학생이 속한 강좌/강의 조회 */
func (c *courseService) GetCourseByStudentID(r *http.Request) ([]model.Courses, error) {
	user, ok := r.Context().Value(middleware.UserContextKey).(*model.User)
	if !ok || user == nil {
		return nil, errors.New("User not found")
	}

	if user.Role != "student" {
		return nil, errors.New("User is not a student")
	}

	enrollments, err := model.FindEnrollmentsByStudentID(user.ID)
	if err != nil {
		return nil, errors.New("Cannot find enrollments")
	}

	var courses []model.Courses
	for _, enrollment := range enrollments {
		course, err := model.FindCourseByCourseID(enrollment.Courses_id)
		if err != nil {
			return nil, errors.New("Cannot find course")
		}
		courses = append(courses, *course)
	}

	return courses, nil
}
