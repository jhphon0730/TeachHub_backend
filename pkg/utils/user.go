package utils

import (
	"errors"
	"crypto/sha256"
	"encoding/hex"

	"image_storage_server/internal/model"
)

// check Valid User Input 
func CheckValidUserInput(user *model.User) error {
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
	
	// algorithm to hash password ( user <go> package SHA256 )
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))

	return nil
}

