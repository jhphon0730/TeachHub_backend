package middleware

import (
	"context"
	"strings"
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/model"
)

type contextKey string

const UserContextKey contextKey = "user"

// JWT 인증 미들웨어
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorization 헤더에서 토큰 추출
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Bearer 토큰 확인
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			http.Error(w, "Invalid authorization token format", http.StatusUnauthorized)
			return
		}

		// 토큰 검증
		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// ID 출력
		user, err := model.FindUserByID(claims.ID)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
		}

		// 사용자 정보를 Context에 저장
		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
