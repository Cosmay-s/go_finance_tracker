package handlers

import "github.com/cosmay-s/go_finance_tracker/internal/services"

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
