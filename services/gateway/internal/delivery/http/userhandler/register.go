package userhandler

import (
	"myapp/api/gen/user"
	"myapp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

// userRegister godoc
// @Summary Register user
// @Description Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body user.RegisterRequest true "Register payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /users/register [post]
func (h Handler) userRegister(c echo.Context) error {
	var req user.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := h.userClient.Register(c.Request().Context(), &req)
	if err != nil {
		resp, code := httpmsg.Error(err)
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusCreated, res)
}
