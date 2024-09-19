# Image Upload 
curl -X POST http://localhost:8080/upload -F 'image=@test2.png'
curl -X GET localhost:8080/read?imageName=test3.png
## Flow 1 
+----------------+                 +----------------+                  +----------------+
|   Frontend     |                 |    Backend     |                  |   Database     |
| (React App)    |                 |   (Go Server)  |                  |                |
+----------------+                 +----------------+                  +----------------+
       |                                |                                      |
       |  1. Folder Creation / File Upload Request                             |
       +--------------------------------->|                                    |
       |                                |                                      |
       |                                |                                      |
       |                                |  2. Handle Request (service)         |
       |                                |------------------------------------> |
       |                                |                                      |
       |                                |    3. Create Folder / Save File       |
       |                                |    to Storage                          |
       |                                |    +---------------------------+       |
       |                                |    |   Update Metadata in DB  |       |
       |                                |    | (Folder/File Name, Path)  |       |
       |                                |    +---------------------------+       |
       |                                |                                      |
       |                                |                                      |
       |                                |  4. Send Metadata and File Info       |
       |                                |<-------------------------------------|
       |                                |                                      |
       |                                |                                      |
       |                                |  5. Return Folder/File List           |
       |                                |                                      |
       |                                |------------------------------------>|
       |                                |                                      |
       |                                |                                      |
       |                                |  6. Display List to User              |
       |                                |                                      |
       +--------------------------------->|                                      |
       |                                |                                      |
+----------------+                 +----------------+                  +----------------+
|   Frontend     |                 |    Backend     |                  |   Database     |
| (React App)    |                 |   (Go Server)  |                  |                |
+----------------+                 +----------------+                  +----------------+

<hr />

# User
* Middleware
* Service
* Handler
* Model 
* utils ( JWT ) 
