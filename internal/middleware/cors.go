package middleware

import (
	"log"
	"net/http"
)

// CORS 미들웨어
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS 헤더 설정
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH, PUT")

		// OPTIONS 요청에 대한 처리
		if r.Method == "OPTIONS" {
			log.Printf("CORS preflight request from %s\n", r.RemoteAddr)
			w.WriteHeader(http.StatusNoContent)
			return
		}
		
		// 다음 핸들러로 전달
		next.ServeHTTP(w, r)
	})
}
