package producthandler

import (
	"myapp/api/gen/product"
	"myapp/pkg/httpmsg"

	"github.com/labstack/echo/v4"

	"net/http"
)

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
