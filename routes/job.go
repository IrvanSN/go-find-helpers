package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/job"
	"github.com/labstack/echo/v4"
)

type JobRouteController struct {
	JobController *controllers.JobController
}

func (r *JobRouteController) InitRoute(e *echo.Echo) {
	e.POST("/v1/job", r.JobController.Create)
}
