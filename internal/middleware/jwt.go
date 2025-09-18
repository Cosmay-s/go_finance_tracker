package middleware

import (
	"strings"

	"github.com/cosmay-s/go_finance_tracker/internal/services"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(authService *services.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(401, "Требуется аутентификация")
			}
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				return echo.NewHTTPError(401, "Неверный формат токена")
			}
			tokenString := parts[1]

			userID, err := authService.ValidateToken(tokenString)
			if err != nil {
				return echo.NewHTTPError(401, "Недействительный токен")
			}

			c.Set("userID", userID)
			return next(c)
		}
	}
}
