package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response 구조체 정의
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSON 응답을 처리하는 함수
func WriteJSONResponse(w http.ResponseWriter, code int, status, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
	}
}

// 성공 응답을 처리하는 함수
func WriteSuccessResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	WriteJSONResponse(w, code, "success", message, data)
}

// 에러 응답을 처리하는 함수
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	WriteJSONResponse(w, code, "error", message, nil)
}
