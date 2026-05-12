package producthandler

import (
	"myapp/api/gen/product"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetProduct godoc
// @Summary Get product
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body product.GetProductByIDRequest true "Get product payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /product/get [get]
func (h Handler) GetProduct(c echo.Context) error {
	var req product.GetProductByIDRequest
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.GetProductByID(c.Request().Context(), &req)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
