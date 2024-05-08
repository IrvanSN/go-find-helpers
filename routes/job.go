package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/job"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/labstack/echo/v4"
)

type JobRouteController struct {
	JobController *controllers.JobController
}

func (r *JobRouteController) InitRoute(e *echo.Echo) {
	e.POST("/v1/callback/xdt", r.JobController.JobPaymentCallback)

	j := e.Group("/v1/jobs")
	j.Use(middlewares.JWTMiddleware)
	j.POST("/post", r.JobController.Create)
	j.POST("/take", r.JobController.Take)
}
