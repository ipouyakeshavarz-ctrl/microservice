package orderhandler

import (
	"myapp/api/gen/order"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) ListUserOrders(c echo.Context) error {
	userID := c.Get("user_id").(uint64)

	resp, err := h.orderClient.ListUserOrders(c.Request().Context(),
		&order.ListUserOrdersRequest{
			UserId: userID,
		})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
