package handlers

import (
	"net/http"
	"image_storage_server/internal/service"
	"image_storage_server/pkg/utils"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	_, err := h.service.RegisterUser(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusCreated, "User registered successfully", nil)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	user, token, err := h.service.LoginUser(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "Login successful", map[string]interface{}{"token": token, "user": user})
	// 위 코드의 any는 interface{}와 같은 의미로 사용된다.
}

