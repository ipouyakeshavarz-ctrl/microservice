package userhandler

import (
	"myapp/api/gen/user"
	"myapp/pkg/httpmsg"
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
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /users/profile [get]
func (h Handler) Profile(c echo.Context) error {

	userID := c.Get("user_id").(uint64)
	resp, err := h.userClient.Profile(
		c.Request().Context(),
		&user.ProfileRequest{UserId: strconv.Itoa(int(userID))},
	)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
