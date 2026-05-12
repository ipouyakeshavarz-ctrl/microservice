package orderhandler

import (
	"myapp/api/gen/order"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetOrder godoc
// @Summary Get order
// @Description Get a specific order by ID
// @Tags Order
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body order.GetOrderRequest true "Get order payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /order/get_order [post]
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
