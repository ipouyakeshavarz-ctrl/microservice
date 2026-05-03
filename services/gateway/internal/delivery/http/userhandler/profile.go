package userhandler

import (
	"myapp/api/gen/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Profile(c echo.Context) error {

	userID := c.Get("user_id").(uint64)
	resp, err := h.userClient.Profile(
		c.Request().Context(),
		&user.ProfileRequest{UserId: strconv.Itoa(int(userID))},
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
