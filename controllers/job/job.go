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

func (jc *JobController) Take(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var jobTake request.JobIdRequest
	if err := c.Bind(&jobTake); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	job, errUseCase := jc.jobUseCase.Take(jobTake.JobIdToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	jobTakeResponse := response.TakeResponseFromUseCase(&job, userData.ID)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success take a Job, please complete the job to get the reward!", jobTakeResponse))
}

func (jc *JobController) JobPaymentCallback(c echo.Context) error {
	var jobPaymentCallback request.JobPaymentCallbackRequest
	if err := c.Bind(&jobPaymentCallback); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	_, errUseCase := jc.jobUseCase.PaymentCallback(jobPaymentCallback.JobPaymentCallbackToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	return nil
}

func (jc *JobController) MarkAsDone(c echo.Context) error {
	var jobDoneRequest request.JobIdRequest
	if err := c.Bind(&jobDoneRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	job, errUseCase := jc.jobUseCase.MarkAsDone(jobDoneRequest.JobIdToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	jobMarkDoneResponse := response.StatusUpdateResponseFromUseCase(&job)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Job completed! All rewards have been credited to the Helpers' account balances. If any Helper slots remain unfulfilled, the corresponding rewards will be refunded to your balance. We appreciate your valuable contribution!", jobMarkDoneResponse))
}

func (jc *JobController) MarkAsOnProgress(c echo.Context) error {
	var jobOnProgressRequest request.JobIdRequest
	if err := c.Bind(&jobOnProgressRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	job, errUseCase := jc.jobUseCase.MarkAsOnProgress(jobOnProgressRequest.JobIdToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	jobOnProgressResponse := response.StatusUpdateResponseFromUseCase(&job)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Job status changed to ON_PROGRESS!", jobOnProgressResponse))
}

func (jc *JobController) GetAllJobs(c echo.Context) error {
	var jobGetAllRequest []entities.Job
	statusFilter := c.QueryParam("status")

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	jobs, errUseCase := jc.jobUseCase.GetAll(&jobGetAllRequest, userData, statusFilter)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	jobGetAllResponse := response.GetAllResponseFromUseCase(&jobs)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all Jobs data!", jobGetAllResponse))
}

func NewJobController(jobUseCase entities.JobUseCaseInterface) *JobController {
	return &JobController{jobUseCase: jobUseCase}
}
