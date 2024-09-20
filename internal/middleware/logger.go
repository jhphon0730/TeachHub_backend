package middleware

import (
	"log"
	"net/http"
)

// 로깅 미들웨어
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 요청을 로깅
		log.Printf("Request: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		
		// 다음 핸들러로 전달
		next.ServeHTTP(w, r)
	})
}

