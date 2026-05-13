package storehandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// updateStore godoc
// @Summary Update store
// @Tags Store
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.UpdateStoreRequest true "Update store payload"
// @Success 200 {object} dto.UpdateStoreResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /store/update [post]
func (h Handler) updateStore(c echo.Context) error {
	var req dto.UpdateStoreRequest
	userID := c.Get("user_id").(uint64)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.storeClient.UpdateStore(c.Request().Context(), &dto.UpdateStoreRequest{
		StoreID:     req.StoreID,
		UserID:      userID,
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
