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
	c := e.Group("/v1/callback")
	c.Use(middlewares.CallbackAuth)
	c.POST("/xdt", r.JobController.JobPaymentCallback)

	j := e.Group("/v1/jobs")
	j.Use(middlewares.JWTMiddleware)
	j.POST("/post", r.JobController.Create)
	j.POST("/take", r.JobController.Take)
	j.POST("/done", r.JobController.MarkAsDone)
	j.POST("/on-progress", r.JobController.MarkAsOnProgress)
}
