package producthandler

import (
	"net/http"
	"productapp/internal/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) updateProduct(c echo.Context) error {
	var req param.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productSvc.Update(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}
