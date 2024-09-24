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
  -d '{"username": "john_doe", "password": "example123", "email": "john@example.com"}'
```
