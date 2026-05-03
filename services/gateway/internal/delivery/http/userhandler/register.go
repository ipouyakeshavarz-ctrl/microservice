package userhandler

import (
	"myapp/api/gen/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) userRegister(c echo.Context) error {
	var req user.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := h.userClient.Register(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}
