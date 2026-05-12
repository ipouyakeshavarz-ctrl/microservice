package orderhandler

import (
	"myapp/api/gen/order"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ListUserOrders godoc
// @Summary List user orders
// @Description Returns all orders for authenticated user
// @Tags Order
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /order/list [get]
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
