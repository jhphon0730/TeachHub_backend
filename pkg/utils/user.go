package utils

import (
	"errors"
	"crypto/sha256"
	"encoding/hex"

	"image_storage_server/internal/model"
)

type LoginUser struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}

// check Valid User Input [Register]
func CheckValidRegisterUserInput(user *model.User) error {
	if user.Email == "" {
		return errors.New("Email is required")
	}
	if user.Username == "" {
		return errors.New("Username is required")
	}
	if user.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

// hash User Password ( When Register User ) 
func HashUserPassword(user *model.User) error {
	if user.Password == "" {
		return errors.New("Password is required")
	}
	
	// algorithm to hash password ( user <go> package hex SHA256 )
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))

	return nil
}

// check Valid User Input [Login]
func CheckValidLoginUserInput(user *LoginUser) error {
	if user.Username == "" {
		return errors.New("Username is required")
	}
	if user.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

// verify User Password ( When Login )
func VerifyUserPassword(inputPassword string, hashedPassword string) error {
	if inputPassword == "" {
		return errors.New("Password is required")
	}

	// Hash the input password using the same algorithm (SHA256)
	hash := sha256.New()
	hash.Write([]byte(inputPassword))
	inputPasswordHashed := hex.EncodeToString(hash.Sum(nil))

	// Compare the hashed input password with the stored hashed password
	if inputPasswordHashed != hashedPassword {
		return errors.New("Invalid password")
	}

	return nil
}
