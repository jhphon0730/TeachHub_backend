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
	user, err := h.service.RegisterUser(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusCreated, "User registered successfully", user)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	token, err := h.service.LoginUser(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "Login successful", map[string]string{"token": token})
}

func (h *UserHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if len(email) == 0 {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Email is required")
		return
	}
	user, err := h.service.FindUserByEmail(email)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "User found", user)
}
