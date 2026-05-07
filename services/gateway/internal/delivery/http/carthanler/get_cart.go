package carthanler

import (
	"myapp/api/gen/cart"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
