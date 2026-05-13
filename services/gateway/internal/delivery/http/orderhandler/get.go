package orderhandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetOrder godoc
// @Summary Get order
// @Description Get a specific order by ID
// @Tags Orders
// @Security BearerAuth
// @Produce json
// @Param id path integer true "Order ID" example(15)
// @Success 200 {object} dto.GetOrderResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /orders/{id} [get]
func (h Handler) GetOrder(c echo.Context) error {
	userID := c.Get("user_id").(uint64)

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid order id")
	}

	resp, gErr := h.orderClient.GetOrder(c.Request().Context(), &dto.GetOrderRequest{
		UserID:  userID,
		OrderID: strconv.FormatUint(orderID, 10),
	},
	)
	if gErr != nil {
		resp, code := httpmsg.Error(gErr)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
