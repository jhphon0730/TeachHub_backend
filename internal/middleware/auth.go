package middleware

// auth ( JWT ) 미들웨어 필수 목록으로는
// 1. 사용자가 있는 지 확인 ( Beer 확인 ) -> Context에 저장 이후 Handler로 보내기
// 2. 관리자 사용자인지 확인해주는 Middleware ( "/admin*" ) 전용 Handler 
	// -> JWT 토큰 유효 확인은 pkg/auth/auth.go 에서 처리


