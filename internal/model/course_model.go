package model

import (
	"time"
	"errors"
)

// 강의 테이블 ( instructor_id<users> )
type Courses struct {
	ID        		int64    	`json:"id"`
	Instructor_id int64    	`json:"instructor_id"` // 1 : N
	Title 				string 		`json:"title"`
	Description 	string 		`json:"description"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}

// 수강 테이블 ( courses_id<courses>, user_id<users> ) 
type Enrollments struct {
	ID 							int64    	`json:"id"`
	Courses_id 			int64    	`json:"courses_id"` // 1 : N
	Student_id 				int64   `json:"student_id"` 		// 1 : N
	Enrollment_date time.Time `json:"enrollment_date"`
}

// ####################### Courses ####################### //
func CreateCoursesTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS courses (
		id INT AUTO_INCREMENT PRIMARY KEY,
		instructor_id INT NOT NULL,
		title VARCHAR(50) NOT NULL,
		description VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

		FOREIGN KEY (instructor_id) REFERENCES users(id) ON DELETE CASCADE
	)
	`
	_, err := DB.Exec(createTableQuery)
	return err
}

// 새로운 강의 추가 
func InsertCourse(courses *Courses) (int64, error) {
	// Check User is instructor
	user, err := FindUserByID(courses.Instructor_id)
	if err != nil {
		return 0, errors.New("Cannot find user")
	}

	if user.Role != "instructor" {
		return 0, errors.New("User is not instructor")
	}

	query := "INSERT INTO courses (instructor_id, title, description) VALUES (?, ?, ?)"

	result, err := DB.Exec(query, courses.Instructor_id, courses.Title, courses.Description)
	if err != nil {
		return 0, err
	}

	coursesID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return coursesID, nil
}

// 모든 강의 조회
func FindAllCourses() ([]Courses, error) {
	query := "SELECT id, instructor_id, title, description, created_at, updated_at FROM courses"

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Courses
	for rows.Next() {
		var course Courses
		if err := rows.Scan(&course.ID, &course.Instructor_id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

// 강사가 생성한 강의들 모두 조회 
func FindCoursesByInstructorID(instructor_id int64) ([]Courses, error) {
	query := "SELECT id, instructor_id, title, description, created_at, updated_at FROM courses WHERE instructor_id = ?"

	rows, err := DB.Query(query, instructor_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Courses
	for rows.Next() {
		var course Courses
		if err := rows.Scan(&course.ID, &course.Instructor_id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

// 강의 ID로 강의 조회
func FindCoursesByCourseID(course_id int64) (*Courses, error) {
	query := "SELECT id, instructor_id, title, description, created_at, updated_at FROM courses WHERE id = ?"

	var courses Courses
	err := DB.QueryRow(query, course_id).Scan(&courses.ID, &courses.Instructor_id, &courses.Title, &courses.Description, &courses.CreatedAt, &courses.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &courses, nil
}

// 수강 내역 테이블에서 조회한 데이터 ( course_id ) 로 강의 조회 
func FindCoursesByEnrollments(enrollments []Enrollments) ([]Courses, error) {
	var courses []Courses
	for _, enrollment := range enrollments {
		course, err := FindCoursesByCourseID(enrollment.Courses_id)
		if err != nil {
			return nil, err
		}
		courses = append(courses, *course)
	}

	return courses, nil
}

// ####################### Enrollments ####################### //
func CreateEnrollmentsTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS enrollments (
		id INT AUTO_INCREMENT PRIMARY KEY,
		courses_id INT NOT NULL,
		student_id INT NOT NULL,
		enrollment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY (courses_id) REFERENCES courses(id),
		FOREIGN KEY (student_id) REFERENCES users(id)
	)
	`
	_, err := DB.Exec(createTableQuery)
	return err
}

// 강의에 새로운 사용자(수강 학생) 추가 
func InsertEnrollmentsByCourseID(course_id int64, student_id int64) (int64, error) {
	query := "INSERT INTO enrollments (courses_id, student_id) VALUES (?, ?)"

	result, err := DB.Exec(query, course_id, student_id)
	if err != nil {
		return 0, err
	}

	enrollmentsID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return enrollmentsID, nil
}

// 학생이 속한 강의 ID들 조회 
func FindEnrollmentsByStudentID(student_id int64) ([]Enrollments, error) {
	query := "SELECT id, courses_id, student_id, enrollment_date FROM enrollments WHERE student_id = ?"

	rows, err := DB.Query(query, student_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []Enrollments
	for rows.Next() {
		var enrollment Enrollments
		if err := rows.Scan(&enrollment.ID, &enrollment.Courses_id, &enrollment.Student_id, &enrollment.Enrollment_date); err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}

