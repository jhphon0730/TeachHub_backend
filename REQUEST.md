# ############################## User ##############################
* [Register](회원가입) OK 
```
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "tester123", "password": "Asdqwe123@", "email": "tester123@gmail.com"}'
```

* [Login](로그인) OK
```
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "tester123", "password": "Asdqwe123@"}'
```

# ############################## Course ##############################
* [Create](강의 생성) OK
```
curl -X POST http://localhost:8080/course -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"title": "인공지능 뿌시기", "description": "인공지능의 시작을 알린 퍼텝스톤"}'
```

* [GetCourseByInstructorID](강의ID로 강의 조회) OK
```
curl -X GET http://localhost:8080/course/instructor?instructor_id=3 -H "Content-Type: application/json" -H "Authorization: Bearer "
```

* [RemoveStudentToCourse](강의ID, 학생ID로 수강취소) OK
```
curl -X DELETE http://localhost:8080/course/student -H "Content-Type: application/json" -H "Authorization: Bearer "
```

* [GetStudentsByCourseID](강의ID로 학생들을 조회) OK
```
curl -X GET http://localhost:8080/course/student?course_id=3 -H "Content-Type: application/json
```

# ############################## Enrollments ##############################
* [AddStudentEnrollment](강의ID, 학생 Usernamed으로 수강신청)
```
curl -X POST http://localhost:8080/enrollment/student -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"student_username": "jhkim123", "course_id": 3}'
```

* [GetCourseByStudentID](학생ID로 수강신청 조회) OK
```
curl -X GET http://localhost:8080/enrollment/student?student_id=2 -H "Content-Type: application" -H "Authorization: Bearer "
```

# ############################## Dashboard ##############################
* [InitialStudentDashboard](사용자 대시보드 초기화) OK
```
curl -X GET http://localhost:8080/dashboard/student/initial -H "Content-Type: application/json" -H "Authorization: Bearer "
```

* [InitialInstructorDashboard](강사 대시보드 초기화) OK
```
curl -X GET http://localhost:8080/dashboard/instructor/initial -H "Content-Type: application/json" -H "Authorization: Bearer "
```

