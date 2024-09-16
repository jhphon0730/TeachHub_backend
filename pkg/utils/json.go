package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response 구조체는 성공 및 에러 응답을 위한 공통 구조를 정의합니다.
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// WriteJSONResponse는 JSON 형식의 응답을 작성합니다.
func writeJSONResponse(w http.ResponseWriter, code int, status, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
	}
}

// WriteSuccessResponse는 성공 응답을 작성합니다.
func WriteSuccessResponse(w http.ResponseWriter, code int, data interface{}) {
	writeJSONResponse(w, code, "success", "", data)
}

// WriteErrorResponse는 에러 응답을 작성합니다.
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	writeJSONResponse(w, code, "error", message, nil)
}
