package handlers

import (
	"net/http"

	"image_storage_server/pkg/utils"
	"image_storage_server/internal/context"
	"image_storage_server/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

func (h *UserHandler) RegisterUser(c *context.Context) {
	user, err := h.service.RegisterUser(c.Request)
	if err != nil {
		utils.WriteErrorResponse(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteSuccessResponse(c.Writer, http.StatusCreated, "User registered successfully", user)
}

func (h *UserHandler) LoginUser(c *context.Context) {
	token, err := h.service.LoginUser(c.Request)
	if err != nil {
		utils.WriteErrorResponse(c.Writer, http.StatusUnauthorized, err.Error())
		return
	}

	utils.WriteSuccessResponse(c.Writer, http.StatusOK, "Login successful", map[string]string{"token": token})
}

func (h *UserHandler) FindUser(c *context.Context) {
	email := c.Request.URL.Query().Get("email")
	if len(email) == 0 {
		utils.WriteErrorResponse(c.Writer, http.StatusBadRequest, "Email is required")
		return
	}

	user, err := h.service.FindUserByEmail(email)
	if err != nil {
		utils.WriteErrorResponse(c.Writer, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteSuccessResponse(w.Writer, http.StatusOK, "User found", user)
}
