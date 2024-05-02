package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/user"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	UserController *controllers.UserController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/signup", r.UserController.SignUp)
}
