package producthandler

import (
	"myapp/api/gen/product"
	"myapp/pkg/httpmsg"

	"github.com/labstack/echo/v4"

	"net/http"
)

// createProduct godoc
// @Summary Create product
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body product.CreateProductRequest true "Create product payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /product/create [post]
func (h Handler) createProduct(c echo.Context) error {

	var req product.CreateProductRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.CreateProduct(c.Request().Context(), &req)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
