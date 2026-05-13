package carthandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
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
// @Param request body dto.AddItemRequest true "Add item payload"
// @Success 200 {object} dto.AddItemResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cart/add_item [post]
func (h Handler) AddItem(c echo.Context) error {
	var req dto.AddItemRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.cartClient.AddItem(c.Request().Context(), &dto.AddItemRequest{
		UserID:    c.Get("user_id").(uint64),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)

}
