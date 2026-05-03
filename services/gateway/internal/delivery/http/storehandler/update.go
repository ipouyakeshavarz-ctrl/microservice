package storehandler

import (
	"myapp/api/gen/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) updateStore(c echo.Context) error {
	var req store.UpdateStoreRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeClient.UpdateStore(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
