package producthandler

import (
	"myapp/api/gen/product"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) deleteProduct(c echo.Context) error {
	var req product.DeleteProductRequest
	userID := c.Get("user_id").(uint64)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.productClient.DeleteProduct(c.Request().Context(), &product.DeleteProductRequest{
		UserId:    userID,
		ProductId: req.ProductId,
		StoreId:   req.StoreId,
	})
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)

}
