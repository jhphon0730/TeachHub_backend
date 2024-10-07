package dto

type UpdateUserDTO struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Skills    []string  `json:"skills"` // skills를 배열로 처리
}
