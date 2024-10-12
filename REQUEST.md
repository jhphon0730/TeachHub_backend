# User

* [Register](회원가입)
```
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "tester123", "password": "Asdqwe123@", "email": "tester123@gmail.com"}'
```

* [Login](로그인)
```
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "tester123", "password": "Asdqwe123@"}'
```

# Course 

## Course
* [Create](강의 생성)
```
curl -X POST http://localhost:8080/course \
    -H "Content-Type: application/json" \
    -d '{"title": "인공지능 뿌시기", "description": "인공지능의 시작을 알린 퍼텝스톤", "instructor_id": 1}'
```

* [FindAllCourses](모든 강의 조회)
```
curl -X GET http://localhost:8080/course \
```

* [FindCoursesByInstructorID](강사의 강의들 조회)
```
curl -X GET http://localhost:8080/course/instructor/1 \
```

* [FindCoursesByCourseID](강의ID로 강의 조회)
```
curl -X GET http://localhost:8080/course/1 \
```

* [FindCoursesByEnrollments](학생이 수강신청한 강의들 조회)
```
curl -X GET http://localhost:8080/course/enrollments/1 \
```

## Enrollments
* [InsertEnrollmentsByCourseID](강의ID, 학생ID로 수강신청)
```
curl -X POST http://localhost:8080/course/enrollments \
  -H "Content-Type: application/json" \
  -d '{"courseID": 1, "studentID": 1}'
```

* [FindEnrollmentsByStudentID](학생ID로 수강신청 조회)
```
curl -X GET http://localhost:8080/course/enrollments/student/1 \
```
