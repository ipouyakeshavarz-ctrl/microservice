package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserClient interface {
	GetProfile(userID string) (interface{}, error)
}

type UserHandler struct {
	user UserClient
}

func NewUserHandler(u UserClient) *UserHandler {
	return &UserHandler{user: u}
}

func (h *UserHandler) Profile(c echo.Context) error {

	userID := c.Get("user_id").(string)

	user, err := h.user.GetProfile(userID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
