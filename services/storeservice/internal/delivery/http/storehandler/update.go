package storehandler

import (
	"net/http"
	"storeapp/internal/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) updateStore(c echo.Context) error {
	var req param.UpdateStoreRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeSvc.UpdateStore(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}
