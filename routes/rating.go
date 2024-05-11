package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/rating"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/labstack/echo/v4"
)

type RatingRouteController struct {
	RatingController *controllers.RatingController
}

func (r *RatingRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/ratings")
	c.Use(middlewares.JWTMiddleware)
	c.POST("", r.RatingController.Create)
	c.GET("", r.RatingController.GetAll)
	c.PUT("/:id", r.RatingController.Update)
	c.DELETE("/:id", r.RatingController.Delete)
}
