package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/user"
	"github.com/labstack/echo/v4"
)

type UserRouteController struct {
	UserController *controllers.UserController
}

func (r *UserRouteController) InitRoute(e *echo.Echo) {
	e.POST("/v1/sign-up", r.UserController.SignUp)
	e.POST("/v1/sign-in", r.UserController.SignIn)
}
