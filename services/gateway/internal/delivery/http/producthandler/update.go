package producthandler

import (
	"myapp/api/gen/product"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) updateProduct(c echo.Context) error {
	var req product.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.UpdateProduct(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}
