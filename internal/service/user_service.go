package service

import (
	"net/http"


	"image_storage_server/pkg/utils"
	"image_storage_server/internal/model"
)

type UserService interface {
	RegisterUser(r *http.Request) (*model.User, error)
	LoginUser(r *http.Request) (string, error)
	FindUserByEmail(email string) (*model.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) RegisterUser(r *http.Request) (*model.User, error) {
	var user model.User
	var err error

	if err = utils.ParseJSON(r, &user); err != nil {
		return nil, err
	}

	// TODO: Validate user input
	// TODO: Check if user already exists
	// TODO: Hash user password

	user.ID, err = model.InsertUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
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
