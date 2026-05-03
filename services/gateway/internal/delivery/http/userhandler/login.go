package userhandler

import (
	"myapp/api/gen/user"

	"github.com/labstack/echo/v4"

	"net/http"
)

func (h Handler) userLogin(c echo.Context) error {
	var req user.LoginRequest
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.userClient.Login(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
