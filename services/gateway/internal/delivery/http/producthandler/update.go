package producthandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// updateProduct godoc
// @Summary Update product
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.UpdateProductRequest true "Update product payload"
// @Success 201 {object} dto.UpdateProductResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /product/update [post]
func (h Handler) updateProduct(c echo.Context) error {
	var req dto.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.UpdateProduct(c.Request().Context(), &req)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusCreated, resp)
}
