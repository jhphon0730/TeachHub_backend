package dto

type CreateCoursesDTO struct {
	Instructor_id int64    	`json:"instructor_id"`
	Title 				string 		`json:"title"`
	Description 	string 		`json:"description"`
}

// 강사가 학생을 강의에 등록할 때 사용하는 DTO
type AddStudentDTO struct {
	Course_id int64 	`json:"course_id"`
	Student_id int64 	`json:"student_id"`
}
