package userhandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Profile godoc
// @Summary Get user profile
// @Description Returns authenticated user profile
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/profile [get]
func (h Handler) Profile(c echo.Context) error {

	userID := c.Get("user_id").(uint64)
	resp, err := h.userClient.Profile(
		c.Request().Context(),
		&dto.ProfileRequest{UserId: strconv.Itoa(int(userID))},
	)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
