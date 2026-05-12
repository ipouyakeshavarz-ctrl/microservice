package storehandler

import (
	"myapp/api/gen/store"
	"myapp/pkg/errmsg"
	"myapp/pkg/httpmsg"

	"github.com/labstack/echo/v4"

	"net/http"
)

// createStore godoc
// @Summary Create store
// @Tags Store
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body store.CreateStoreRequest true "Create store payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Router /store/create [post]
func (h Handler) createStore(c echo.Context) error {
	var req store.CreateStoreRequest
	userID := c.Get("user_id").(uint64)
	if req.UserId == userID {
		return echo.NewHTTPError(http.StatusForbidden, errmsg.ErrorMsgYouAreNotAuthorizedToCreateAStore)
	}
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeClient.CreateStore(c.Request().Context(), &store.CreateStoreRequest{
		UserId:      userID,
		Name:        req.Name,
		Description: req.Description,
		PhoneNumber: req.PhoneNumber,
		LogoUrl:     req.LogoUrl,
		Address:     req.Address,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
