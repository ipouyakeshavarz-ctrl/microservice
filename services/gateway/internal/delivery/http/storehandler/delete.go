package storehandler

import (
	"myapp/api/gen/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) deleteStore(c echo.Context) error {
	var req store.DeleteStoreRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeClient.DeleteStore(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, resp)

}
