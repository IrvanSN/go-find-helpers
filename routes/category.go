package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/category"
	"github.com/labstack/echo/v4"
)

type CategoryRouteController struct {
	CategoryController *controllers.CategoryController
}

func (r *CategoryRouteController) InitRoute(e *echo.Echo) {
	e.POST("/v1/category", r.CategoryController.Create)
}
