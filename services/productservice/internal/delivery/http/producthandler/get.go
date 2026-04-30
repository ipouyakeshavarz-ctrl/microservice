package producthandler

import (
	"myapp/pkg/httpmsg"
	"net/http"
	"productapp/internal/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetProduct(c echo.Context) error {
	var req param.GetProductByIDRequest
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}
	resp, err := h.productSvc.GetByID(c.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
