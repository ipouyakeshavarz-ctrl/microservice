package producthandler

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/httpmsg"

	"github.com/labstack/echo/v4"

	"net/http"
)

// createProduct godoc
// @Summary Create product
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequest true "Create product payload"
// @Success 200 {object} dto.CreateProductResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /product/create [post]
func (h Handler) createProduct(c echo.Context) error {

	var req dto.CreateProductRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.CreateProduct(c.Request().Context(), &req)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
