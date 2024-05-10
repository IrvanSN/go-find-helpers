package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/user"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/labstack/echo/v4"
)

type UserRouteController struct {
	UserController *controllers.UserController
}

func (r *UserRouteController) InitRoute(e *echo.Echo) {
	e.POST("/v1/sign-up", r.UserController.SignUp)
	e.POST("/v1/sign-in", r.UserController.SignIn)

	u := e.Group("/v1/users")
	u.Use(middlewares.JWTMiddleware)
	u.POST("/addresses", r.UserController.AddAddress)
	u.GET("/addresses", r.UserController.GetAllAddresses)
	u.PUT("/:id", r.UserController.Update)
	u.GET("/:id", r.UserController.Find)
	u.DELETE("/:id", r.UserController.Delete)
}
