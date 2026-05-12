package userhandler

import (
	"myapp/api/gen/user"
	"myapp/pkg/httpmsg"

	"net/http"

	"github.com/labstack/echo/v4"
)

// userLogin godoc
// @Summary User login
// @Description Login user with phone number and password
// @Tags Users
// @Accept json
// @Produce json
// @Param request body user.LoginRequest true "Login payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /users/login [post]
func (h Handler) userLogin(c echo.Context) error {
	var req user.LoginRequest
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.userClient.Login(c.Request().Context(), &req)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
