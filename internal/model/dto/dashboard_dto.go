package dto 

type InitialDashboardDTO struct {
	TotalStudentCount int `json:"total_student_count"`
	TotalInstructorCount int `json:"total_instructor_count"`
	TotalCourseCount int `json:"total_course_count"`
	MyCourseCount int `json:"my_course_count"`
}

type InitialStudentDashboardDTO struct {
	// Total Student Count
	// Total Instructor Count
	// Total Course Count
	InitialDashboardDTO
}

type InitialInstructorDashboardDTO struct {
	// Total Student Count
	// Total Instructor Count
	// Total Course Count
	// My Course Count
	InitialDashboardDTO
}
