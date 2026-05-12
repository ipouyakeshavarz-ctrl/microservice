package producthandler

import (
	"myapp/api/gen/product"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// updateProduct godoc
// @Summary Update product
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body product.UpdateProductRequest true "Update product payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /product/update [post]
func (h Handler) updateProduct(c echo.Context) error {
	var req product.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.UpdateProduct(c.Request().Context(), &req)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusCreated, resp)
}
