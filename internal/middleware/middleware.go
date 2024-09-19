package middleware

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

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

// 로깅 미들웨어
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 요청을 로깅
		log.Printf("Request: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		
		// 다음 핸들러로 전달
		next.ServeHTTP(w, r)
	})
}

// 미들웨어 체인을 생성하는 함수
func ChainMiddleware(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, x := range xs {
			next = x(next)
		}
		return next
	}
}
