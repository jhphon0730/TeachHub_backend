package dto

import (
	"time"
)

// 강사가 강좌/강의를 생성할 때 사용하는 DTO
type CreateCoursesDTO struct {
	Instructor_id int64    	`json:"instructor_id"`
	Title 				string 		`json:"title"`
	Description 	string 		`json:"description"`
}

// 강사가 소속된 강좌/강의를 조회할 때 사용하는 DTO
type FindCourseByInstructorIDDTO struct { 
	ID        		int64    	`json:"id"`
	Instructor_id int64    	`json:"instructor_id"` // 1 : N
	Title 				string 		`json:"title"`
	Description 	string 		`json:"description"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
	StrudentCount int64 		`json:"student_count"`
}

// 학생이 속한 강좌/강의를 조회할 때 사용하는 DTO
type RemoveStudentDTO struct {
	Course_id int64 	`json:"course_id"`
	Student_Username string 	`json:"student_username"`
}

type FindStudentsByCourseIDDTO struct {
	User_id int64 	`json:"user_id"`
	User_Username string 	`json:"user_username"`
}
