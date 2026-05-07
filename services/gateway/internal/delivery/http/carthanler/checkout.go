package carthanler

import (
	"myapp/api/gen/cart"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) Checkout(c echo.Context) error {

	resp, err := h.cartClient.Checkout(c.Request().Context(), &cart.CheckoutRequest{
		UserId: c.Get("user_id").(uint64),
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)

}
