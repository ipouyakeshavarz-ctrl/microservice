package storehandler

import (
	"net/http"
	"storeapp/internal/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) deleteStore(c echo.Context) error {
	var req param.DeleteStoreRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := h.storeSvc.DeleteStore(c.Request().Context(), req)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

}
