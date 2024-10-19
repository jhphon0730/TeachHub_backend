package model

import (
	"time"
	"errors"

	"image_storage_server/internal/model/dto"
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
func FindCourseByInstructorID(instructor_id int64) ([]dto.FindCourseByInstructorIDDTO, error) {
	query := `
		SELECT c.*, IFNULL(COUNT(e.id), 0) AS students_count
		FROM courses c
		LEFT JOIN enrollments e ON c.id = e.courses_id
		WHERE c.instructor_id = ?
		GROUP BY c.id;
	`

	rows, err := DB.Query(query, instructor_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []dto.FindCourseByInstructorIDDTO
	for rows.Next() {
		var course dto.FindCourseByInstructorIDDTO
		if err := rows.Scan(&course.ID, &course.Instructor_id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.StrudentCount); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

// 강의 ID로 강의 조회
func FindCourseByCourseID(course_id int64) (*dto.FindCourseByInstructorIDDTO, error) {
	query := `
		SELECT c.*, IFNULL(COUNT(e.id), 0) AS students_count
		FROM courses c
		LEFT JOIN enrollments e ON c.id = e.courses_id
		WHERE c.id = ?
		GROUP BY c.id;
	`

	var courses dto.FindCourseByInstructorIDDTO
	err := DB.QueryRow(query, course_id).Scan(&courses.ID, &courses.Instructor_id, &courses.Title, &courses.Description, &courses.CreatedAt, &courses.UpdatedAt, &courses.StrudentCount)
	if err != nil {
		return nil, err
	}

	return &courses, nil
}

// 수강 내역 테이블에서 조회한 데이터 ( course_id ) 로 강의 조회 
func FindCourseByEnrollments(enrollments []Enrollments) ([]dto.FindCourseByInstructorIDDTO, error) {
	var courses []dto.FindCourseByInstructorIDDTO
	for _, enrollment := range enrollments {
		course, err := FindCourseByCourseID(enrollment.Courses_id)
		if err != nil {
			return nil, err
		}
		courses = append(courses, *course)
	}

	return courses, nil
}

