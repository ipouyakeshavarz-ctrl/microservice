package carthandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCart godoc
// @Summary Get cart
// @Description Returns authenticated user's cart
// @Tags Cart
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.GetCartResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cart/get_cart [get]
func (h Handler) GetCart(c echo.Context) error {
	userID := c.Get("user_id").(uint64)
	resp, err := h.cartClient.GetCart(c.Request().Context(), &dto.GetCartRequest{
		UserID: userID,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
