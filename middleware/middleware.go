package middleware

import (
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		// Continue to the next handler
		next.ServeHTTP(w, r)
	})
}
