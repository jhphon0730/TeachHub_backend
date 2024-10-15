package service

import (
	"errors"
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/model"
	"image_storage_server/internal/model/dto"
	"image_storage_server/internal/middleware"
)

type EnrollmentService interface {
	AddStudentEnrollment(r *http.Request) (error)
	GetCourseByStudentID(r *http.Request) ([]model.Courses, error)
}

type enrollmentService struct { }

func NewEnrollmentService() EnrollmentService {
	return &enrollmentService{}
}

/* 강사가 학생의 ID로 수강 내역 추가 */
func (c *enrollmentService) AddStudentEnrollment(r *http.Request) error {
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

	if addStudentDTO.Course_id == 0 || addStudentDTO.Student_Username == "" {
		return errors.New("Course_id or Student_id is empty")
	}

	student, err := model.FindUserByUserName(addStudentDTO.Student_Username)
	if err != nil {
		return errors.New("Cannot find student")
	}
	_, err = model.InsertStudentEnrollment(addStudentDTO.Course_id, student.ID)
	if err != nil {
		return errors.New("Cannot add student to course")
	}

	return nil
}

/* 학생 ID로 학생이 속한 강좌/강의 조회 */
func (c *enrollmentService) GetCourseByStudentID(r *http.Request) ([]model.Courses, error) {
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
