package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// 미들웨어 체인을 생성하는 함수
func ChainMiddleware(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, x := range xs {
			next = x(next)
		}
		return next
	}
}

