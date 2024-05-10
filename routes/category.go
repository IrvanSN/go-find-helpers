package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/category"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/labstack/echo/v4"
)

type CategoryRouteController struct {
	CategoryController *controllers.CategoryController
}

func (r *CategoryRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/categories")
	c.Use(middlewares.JWTMiddleware)
	c.POST("", r.CategoryController.Create)
	c.PUT("/:id", r.CategoryController.Update)
	c.DELETE("/:id", r.CategoryController.Delete)
	c.GET("", r.CategoryController.GetAll)
}
