package storehandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"myapp/pkg/errmsg"

	"github.com/labstack/echo/v4"

	"net/http"
)

// createStore godoc
// @Summary Create store
// @Tags Store
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateStoreRequest true "Create store payload"
// @Success 200 {object} dto.CreateStoreResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /store/create [post]
func (h Handler) createStore(c echo.Context) error {
	var req dto.CreateStoreRequest
	userID := c.Get("user_id").(uint64)
	if req.UserId == userID {
		return echo.NewHTTPError(http.StatusForbidden, errmsg.ErrorMsgYouAreNotAuthorizedToCreateAStore)
	}
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeClient.CreateStore(c.Request().Context(),&req)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
