package producthandler

import (
	"github.com/labstack/echo/v4"

	"net/http"
	"productapp/internal/param"
)

func (h Handler) createProduct(c echo.Context) error {
	var req param.CreateProductRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productSvc.Create(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
