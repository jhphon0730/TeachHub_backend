package service

import (
	"net/http"

	"image_storage_server/internal/model"
)

type UserService interface {
	RegisterUser(r *http.Request) (*model.User, error)
	LoginUser(r *http.Request) (string, error)  // Returns a token
	FindUserByEmail(email string) (*model.User, error)
}

type userService struct {
	// Add any necessary fields, like a database connection
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) RegisterUser(r *http.Request) (*model.User, error) {
	// Extract user details from request
	// Check if user already exists
	// Hash password and save user to database
	// Return user and/or error
	return &model.User{}, nil
}

func (s *userService) LoginUser(r *http.Request) (string, error) {
	// Extract user credentials from request
	// Verify user credentials
	// Generate and return JWT token
	return "", nil
}

func (s *userService) FindUserByEmail(email string) (*model.User, error) {
	// Find and return user by email
	return &model.User{}, nil
}
