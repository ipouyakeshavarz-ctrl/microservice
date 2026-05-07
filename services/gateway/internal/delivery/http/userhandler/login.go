package userhandler

import (
	"myapp/api/gen/user"
	"myapp/pkg/httpmsg"

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
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
