package userhandler

import (
	"fmt"
	"myapp/api/gen/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) userRegister(c echo.Context) error {
	var req user.RegisterRequest

	fmt.Println("register user")

	if err := c.Bind(&req); err != nil {
		return err
	}
	fmt.Printf("Name: %v | Phone: %v | Pass: %v\n",
		req.Name, req.PhoneNumber, req.Password)

	res, err := h.userClient.Register(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}
