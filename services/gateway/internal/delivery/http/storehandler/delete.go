package storehandler

import (
	"myapp/api/gen/store"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// deleteStore godoc
// @Summary Delete store
// @Tags Store
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body store.DeleteStoreRequest true "Delete store payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /store/delete [post]
func (h Handler) deleteStore(c echo.Context) error {
	var req store.DeleteStoreRequest
	userID := c.Get("user_id").(uint64)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeClient.DeleteStore(c.Request().Context(), &store.DeleteStoreRequest{
		UserId:  userID,
		StoreId: req.StoreId,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, resp)

}
