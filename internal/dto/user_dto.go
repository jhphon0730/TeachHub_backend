package dto 

type RegisterUserDTO struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type LoginUserDTO struct {
	ID  			string 		`json:"id"` // Username || Email ( Contains the value of Username or Email )
	Password  string    `json:"password"`
}
