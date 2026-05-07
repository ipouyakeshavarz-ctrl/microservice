package carthanler

import (
	"myapp/api/gen/cart"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) AddItem(c echo.Context) error {
	var req cart.AddItemRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.cartClient.AddItem(c.Request().Context(), &cart.AddItemRequest{
		UserId:    c.Get("user_id").(uint64),
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)

}
