package producthandler

import (
	"myapp/api/gen/product"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) deleteProduct(c echo.Context) error {
	var req product.DeleteProductRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.DeleteProduct(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}
