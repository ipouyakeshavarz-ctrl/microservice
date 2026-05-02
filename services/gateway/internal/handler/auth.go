package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthClient interface {
	Login(email, password string) (string, error)
}

type AuthHandler struct {
	auth AuthClient
}

func NewAuthHandler(a AuthClient) *AuthHandler {
	return &AuthHandler{auth: a}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c echo.Context) error {

	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	token, err := h.auth.Login(req.Email, req.Password)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
