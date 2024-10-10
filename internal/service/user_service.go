package service

import (
	"net/http"
	"errors"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/model"
	"image_storage_server/internal/model/dto"
)

type UserService interface {
	RegisterUser(r *http.Request) (*model.User, error)
	LoginUser(r *http.Request) (*model.User, string, error)
	UpdateUser(r *http.Request) error
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

	// Validate user input
	if err = utils.CheckValidRegisterUserInput(&user); err != nil {
		return nil, err
	}
	// Check if user already exists
	if alerady_created, err := model.FindUserByUserName(user.Username); err == nil && alerady_created.Username != "" {
		return nil, errors.New("User already exists")
	}
	// Hash user password
	if err = utils.HashUserPassword(&user); err != nil {
		return nil, err
	}

	user.ID, err = model.InsertUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// return JWT token, is error 
func (s *userService) LoginUser(r *http.Request) (*model.User, string, error) {
	var login utils.LoginUser
	var err error

	if err = utils.ParseJSON(r, &login); err != nil {
		return nil, "", err
	}

	// Validate user input
	if err = utils.CheckValidLoginUserInput(&login); err != nil {
		return nil, "", err
	}

	// Find User
	user, err := model.FindUserByUserName(login.Username)
	if err != nil {
		return nil, "", err
	}

	// Decode password
	err = utils.VerifyUserPassword(login.Password, user.Password)
	if err != nil {
		return nil, "", errors.New("Invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// return error 
func (s *userService) UpdateUser(r *http.Request) error {
	var update_user_dto dto.UpdateUserDTO
	var err error

	if err = utils.ParseJSON(r, &update_user_dto); err != nil {
		return err
	}

	// Validate user input
	if err = utils.CheckValidUpdateUserInput(&update_user_dto); err != nil {
		return err
	}

	// Find User
	user, err := model.FindUserByUserName(update_user_dto.Username)
	if err != nil {
		return err
	}

	// Check Valid User Password, Input Password
	// Decode password
	err = utils.VerifyUserPassword(update_user_dto.Password, user.Password)
	if err != nil {
		return errors.New("Invalid password")
	}

	// Update User
	user.Username = update_user_dto.Username
	user.Email = update_user_dto.Email
	user.Bio = update_user_dto.Bio

	err = model.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
