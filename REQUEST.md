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
curl -X POST http://localhost:8080/course -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"title": "인공지능 뿌시기", "description": "인공지능의 시작을 알린 퍼텝스톤"}'
```

* [GetCourseByInstructorID](강의ID로 강의 조회)
```
curl -X GET http://localhost:8080/course?instructor_id=3 -H "Content-Type: application/json" -H "Authorization: Bearer "
```

* [FindAllCourses](모든 강의 조회)
```
```

* [FindCoursesByInstructorID](강사의 강의들 조회)
```
```

* [FindCoursesByEnrollments](학생이 수강신청한 강의들 조회)
```
```

## Enrollments
* [InsertEnrollmentsByCourseID](강의ID, 학생ID로 수강신청)
```
```

* [FindEnrollmentsByStudentID](학생ID로 수강신청 조회)
```
```

# Dashboard
* [InitialStudentDashboard](사용자 대시보드 초기화)
```
curl -X GET http://localhost:8080/dashboard/student/initial -H "Content-Type: application/json" -H "Authorization: Bearer "
```

* [InitialInstructorDashboard](강사 대시보드 초기화)
```
curl -X GET http://localhost:8080/dashboard/instructor/initial -H "Content-Type: application/json" -H "Authorization: Bearer "
```

