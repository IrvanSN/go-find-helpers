package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/address"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/labstack/echo/v4"
)

type AddressRouteController struct {
	AddressController *controllers.AddressController
}

func (r *AddressRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/addresses")
	c.Use(middlewares.JWTMiddleware)
	c.PUT("/:id", r.AddressController.Update)
	c.DELETE("/:id", r.AddressController.Delete)
}
