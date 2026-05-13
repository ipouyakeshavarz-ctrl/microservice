package carthandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Checkout godoc
// @Summary Checkout cart
// @Description Checkout authenticated user's cart
// @Tags Cart
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.CheckoutResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cart/check_out [get]
func (h Handler) Checkout(c echo.Context) error {

	resp, err := h.cartClient.Checkout(c.Request().Context(), &dto.CheckoutRequest{
		UserID: c.Get("user_id").(uint64),
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)

}
