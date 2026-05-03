package middleware

import (
	"net/http"
	"strings"

	"myapp/api/gen/auth"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authClient auth.AuthServiceClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// گرفتن توکن از هدر Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "توکن یافت نشد")
			}
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// call service Auth
			resp, err := authClient.VerifyToken(c.Request().Context(), &auth.VerifyTokenRequest{
				Token: token,
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
			}

			c.Set("user_id", resp.UserId)
			c.Set("role", resp.Role)

			return next(c)
		}
	}
}
