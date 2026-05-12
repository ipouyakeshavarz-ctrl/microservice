package carthanler

import (
	"myapp/api/gen/cart"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCart godoc
// @Summary Get cart
// @Description Returns authenticated user's cart
// @Tags Cart
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /cart/get_cart [get]
func (h Handler) GetCart(c echo.Context) error {
	userID := c.Get("user_id").(uint64)
	resp, err := h.cartClient.GetCart(c.Request().Context(), &cart.GetCartRequest{
		UserId: userID,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
