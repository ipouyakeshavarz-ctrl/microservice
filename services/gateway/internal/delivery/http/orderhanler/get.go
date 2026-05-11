package orderhandler

import (
	"myapp/api/gen/order"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetOrder(c echo.Context) error {
	var req order.GetOrderRequest

	userID := c.Get("user_id").(uint64)

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.orderClient.GetOrder(c.Request().Context(), &order.GetOrderRequest{
		UserId:  userID,
		OrderId: req.OrderId,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
