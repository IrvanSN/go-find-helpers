package controllers

import (
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/labstack/echo/v4"
)

type JobController struct {
	jobUseCase entities.JobUseCaseInterface
}

func (jc *JobController) Create(c echo.Context) error {
	return nil
}

func NewJobController(jobUseCase entities.JobUseCaseInterface) *JobController {
	return &JobController{jobUseCase: jobUseCase}
}
