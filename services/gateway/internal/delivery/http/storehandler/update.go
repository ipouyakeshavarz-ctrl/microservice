package storehandler

import (
	"myapp/api/gen/store"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// updateStore godoc
// @Summary Update store
// @Tags Store
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body store.UpdateStoreRequest true "Update store payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /store/update [post]
func (h Handler) updateStore(c echo.Context) error {
	var req store.UpdateStoreRequest
	userID := c.Get("user_id").(uint64)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeClient.UpdateStore(c.Request().Context(), &store.UpdateStoreRequest{
		StoreId:     req.StoreId,
		UserId:      userID,
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		IsActive:    req.IsActive,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
