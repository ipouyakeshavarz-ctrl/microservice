package storehandler

import (
	"myapp/api/gen/store"

	"github.com/labstack/echo/v4"

	"net/http"
)

func (h Handler) createStore(c echo.Context) error {
	var req store.CreateStoreRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeClient.CreateStore(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
