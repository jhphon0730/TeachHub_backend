package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON 응답을 처리하는 함수
func WriteJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
	}
}

// 에러 응답을 처리하는 함수
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	WriteJSONResponse(w, code, map[string]string{"error": message})
}

