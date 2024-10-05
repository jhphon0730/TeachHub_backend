# Image

* [Upload](#upload)
```
curl -X POST http://localhost:8080/upload -F 'image=@test2.png'
```

* [Read](#read)
```
curl -X GET localhost:8080/read?imageName=test3.png
```

# User

* [Register](#register)
```
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "tester123", "password": "Asdqwe123@", "email": "tester123@gmail.com"}'
```

* [Login](#login)
```
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "tester123", "password": "Asdqwe123@"}'
```
