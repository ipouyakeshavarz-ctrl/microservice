package storehandler

import (
	"myapp/pkg/httpmsg"
	"net/http"
	"storeapp/internal/claim"
	"storeapp/internal/param"

	"github.com/labstack/echo/v4"
)

func (h Handler) listStore(c echo.Context) error {
	claims := claim.GetClaimsFromEchoContext(c)

	resp, err := h.storeSvc.ListStoresByUser(c.Request().Context(),
		param.ListStoresByUserRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
