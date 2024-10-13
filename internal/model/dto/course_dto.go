package dto

type CreateCoursesDTO struct {
	Instructor_id int64    	`json:"instructor_id"`
	Title 				string 		`json:"title"`
	Description 	string 		`json:"description"`
}

