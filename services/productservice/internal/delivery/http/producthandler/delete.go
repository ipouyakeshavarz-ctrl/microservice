package producthandler

import (
	"net/http"
	"productapp/internal/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) deleteProduct(c echo.Context) error {
	var req param.DeleteProductRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := h.productSvc.Delete(c.Request().Context(), req)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

}
