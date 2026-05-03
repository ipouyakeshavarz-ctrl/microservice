package storehandler

import (
	"myapp/api/gen/store"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) listStore(c echo.Context) error {
	userID := c.Get("user_id").(uint64)

	resp, err := h.storeClient.ListStoresByUser(c.Request().Context(),
		&store.ListStoresByUserRequest{UserId: userID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
