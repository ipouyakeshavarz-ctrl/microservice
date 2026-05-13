package producthandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetProduct godoc
// @Summary Get product
// @Description Retrieve product by IDs.
// @Description Example: /product/get?product_id=12&store_id=3
// @Tags Product
// @Security BearerAuth
// @Produce json
// @Param product_id query integer true "Product ID" example(12)
// @Param store_id query integer true "Store ID" example(3)
// @Success 200 {object} dto.GetProductResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /product/get [get]
func (h Handler) GetProduct(c echo.Context) error {
	var req dto.GetProductByIDRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.GetProductByID(
		c.Request().Context(),
		&req,
	)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
