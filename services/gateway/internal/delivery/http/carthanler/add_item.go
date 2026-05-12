package carthanler

import (
	"myapp/api/gen/cart"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AddItem godoc
// @Summary Add item to cart
// @Description Add product to authenticated user's cart
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body cart.AddItemRequest true "Add item payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /cart/add_item [post]
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
