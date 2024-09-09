package main 

import (
	"log"
	"net/http"
	"os"
	"io"
	"image_storage_server/json"
)

const (
	PORT = "0.0.0.0:8080"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		// Continue to the next handler
		next.ServeHTTP(w, r)
	})
}

func SaveImage(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check the File is Already Exists
	//	if the file is already exists, return file path 
	if _, err := os.Stat(handler.Filename); err == nil {
		log.Println("File is already exists")
		json.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"url": "http://localhost:8080/images/" + handler.Filename,
		})
		return
	}

	// Create a new file in the server
	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file to the server
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"url": "http://localhost:8080/images/" + handler.Filename,
	})
}

// 저장 된 이미지를 읽기 위한 함수 
// http://localhost:8080/images/temp_image.png 로 접근 시에 이미지가 떠야 함
func ReadImage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/upload", SaveImage)
	router.HandleFunc("/", ReadImage)

	server := &http.Server{
		Addr:    PORT,
		Handler: CORS(router),
	}
	// 이후에 CORS와 같은 미들웨어를 2개 이상 사용할 때는 순서에 주의해야 함
	// CORS(router) -> Logger(CORS(router)) 이런식으로 사용해야 함

	log.Println("Server is running on port", PORT)
	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
