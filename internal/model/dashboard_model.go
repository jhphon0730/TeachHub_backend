package model

// 모든 학생 수를 반환하는 함수
func GetAllStudentsCount() (int, error) {
	// Query, Count Function, role = student
	query := "SELECT COUNT(*) FROM users WHERE role = 'student'"

	var count int
	err := DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 모든 강사 수를 반환하는 함수
func GetAllInstructorsCount() (int, error) {
	// Query, Count Function, role = instructor
	query := "SELECT COUNT(*) FROM users WHERE role = 'instructor'"

	var count int
	err := DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 모든 코스 수를 반환하는 함수
func GetAllCoursesCount() (int, error) {
	// Query, Count Function
	query := "SELECT COUNT(*) FROM courses"

	var count int
	err := DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 강사가 가지 있는 코스 수를 반환하는 함수
func GetAllMyCoursesCountByInstructorID(instructor_id int64) (int, error) {
	// Query, Count Function
	query := "SELECT COUNT(*) FROM courses WHERE instructor_id = ?"

	var count int
	err := DB.QueryRow(query, instructor_id).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
