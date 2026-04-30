package claim

import (
	"storeapp/internal/auth"
	cfg "storeapp/internal/config"

	"github.com/labstack/echo/v4"
)

func GetClaimsFromEchoContext(c echo.Context) *authservice.Claims {
	return c.Get(cfg.AuthMiddlewareContextKey).(*authservice.Claims)
}
