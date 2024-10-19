package dto

import (
	"time"
)

type CreateCoursesDTO struct {
	Instructor_id int64    	`json:"instructor_id"`
	Title 				string 		`json:"title"`
	Description 	string 		`json:"description"`
}

type FindCourseByInstructorIDDTO struct { 
	ID        		int64    	`json:"id"`
	Instructor_id int64    	`json:"instructor_id"` // 1 : N
	Title 				string 		`json:"title"`
	Description 	string 		`json:"description"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
	StrudentCount int64 		`json:"student_count"`
}
