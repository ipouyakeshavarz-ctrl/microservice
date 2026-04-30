package storehandler

import (
	"myapp/pkg/httpmsg"

	"github.com/labstack/echo/v4"

	"net/http"
	"storeapp/internal/param"
)

func (h Handler) createStore(c echo.Context) error {
	var req param.CreateStoreRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if fieldErrors, err := h.storeValidator.ValidateCreateRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	resp, err := h.storeSvc.CreateStore(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
