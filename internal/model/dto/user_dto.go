package dto

type UpdateUserDTO struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Password string 		`json:"password"`
}
