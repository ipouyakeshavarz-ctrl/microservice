package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	ValidateToken(token string) (string, error)
}

type AuthMiddleware struct {
	auth AuthService
}

func NewAuthMiddleware(a AuthService) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			userID, err := a.ValidateToken(token)

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			c.Set("user_id", userID)

			return next(c)
		}
	}
}
