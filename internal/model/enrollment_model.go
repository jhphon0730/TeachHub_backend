package model

import (
	"time"
)

// 수강 테이블 ( courses_id<courses>, user_id<users> ) 
type Enrollments struct {
	ID 							int64    	`json:"id"`
	Courses_id 			int64    	`json:"courses_id"` // 1 : N
	Student_id 				int64   `json:"student_id"` 		// 1 : N
	Enrollment_date time.Time `json:"enrollment_date"`
}

// ####################### Enrollments ####################### //
func CreateEnrollmentsTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS enrollments (
		id INT AUTO_INCREMENT PRIMARY KEY,
		courses_id INT NOT NULL,
		student_id INT NOT NULL,
		enrollment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY (courses_id) REFERENCES courses(id) ON DELETE CASCADE,
		FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE
	)
	`
	_, err := DB.Exec(createTableQuery)
	return err
}

// 강의에 새로운 사용자(수강 학생) 추가 
func InsertStudentEnrollment(course_id int64, student_id int64) (int64, error) {
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

// 학생 ID와 강의 ID로 수강 정보 조회
func FindEnrollmentByStudentIDAndCourseID(student_id int64, course_id int64) (*Enrollments, error) {
	query := "SELECT id, courses_id, student_id, enrollment_date FROM enrollments WHERE student_id = ? AND courses_id = ?"

	row := DB.QueryRow(query, student_id, course_id)

	var enrollment Enrollments
	if err := row.Scan(&enrollment.ID, &enrollment.Courses_id, &enrollment.Student_id, &enrollment.Enrollment_date); err != nil {
		return nil, err
	}

	return &enrollment, nil
}

// 학생 ID와 강의 ID로 수강 정보 삭제
func DeleteEnrollmentByStudentIDAndCourseID(student_id int64, course_id int64) error {
	query := "DELETE FROM enrollments WHERE student_id = ? AND courses_id = ?"

	_, err := DB.Exec(query, student_id, course_id)
	if err != nil {
		return err
	}

	return nil
}
