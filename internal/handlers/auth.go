package handlers

import (
	"net/http"

	"github.com/cosmay-s/go_finance_tracker/internal/services"
	"github.com/labstack/echo/v4"
)

type AuthHandles struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandles {
	return &AuthHandles{authService: authService}
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandles) Register(c echo.Context) error {
	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"ошибка": err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"ошибка": err.Error(),
		})
	}
	err := h.authService.Register(services.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"ошибка": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"уведомление": "Пользователь создан",
	})
}
