package userhandler

import (
	"myapp/pkg/httpmsg"
	"net/http"
	"strconv"
	"userapp/internal/domain"
	"userapp/internal/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) userProfile(c echo.Context) error {
	claims := GetClaimsFromEchoContext(c)

	resp, err := h.userSvc.Profile(c.Request().Context(),
		param.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
func GetClaimsFromEchoContext(c echo.Context) domain.Claims {
	// خواندن از هدرهایی که Gateway فرستاده
	userIDStr := c.Request().Header.Get("X-User-Id")
	roleStr := c.Request().Header.Get("X-User-Role")

	userID, _ := strconv.Atoi(userIDStr)

	return domain.Claims{
		UserID: uint(userID),
		Role:   domain.MapToRoleEntity(roleStr),
	}
}
