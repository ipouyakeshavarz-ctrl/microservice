package storehandler

import (
	"myapp/api/gen/store"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// listStore godoc
// @Summary List stores
// @Tags Store
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /store/list [get]
func (h Handler) listStore(c echo.Context) error {
	userID := c.Get("user_id").(uint64)

	resp, err := h.storeClient.ListStoresByUser(c.Request().Context(),
		&store.ListStoresByUserRequest{UserId: userID})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
