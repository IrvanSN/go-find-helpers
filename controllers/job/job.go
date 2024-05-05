package controllers

import (
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/job/request"
	"github.com/irvansn/go-find-helpers/controllers/job/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type JobController struct {
	jobUseCase entities.JobUseCaseInterface
}

func (jc *JobController) Create(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var jobCreate request.JobCreateRequest
	if err := c.Bind(&jobCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	job, errUseCase := jc.jobUseCase.Create(jobCreate.JobCreateToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	jobCreateResponse := response.CreateResponseFromUseCase(&job)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success create Job, please pay the commission before the due date!", jobCreateResponse))
}

func NewJobController(jobUseCase entities.JobUseCaseInterface) *JobController {
	return &JobController{jobUseCase: jobUseCase}
}
