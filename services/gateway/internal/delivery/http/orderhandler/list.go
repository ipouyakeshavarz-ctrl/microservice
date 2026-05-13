package orderhandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ListUserOrders godoc
// @Summary List user orders
// @Description Returns all orders for authenticated user
// @Tags Order
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.ListUserOrdersResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order/list [get]
func (h Handler) ListUserOrders(c echo.Context) error {
	userID := c.Get("user_id").(uint64)

	resp, err := h.orderClient.ListUserOrders(c.Request().Context(),
		&dto.ListUserOrdersRequest{
			UserID: userID,
		})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
