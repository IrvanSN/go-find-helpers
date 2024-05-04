package controllers

import (
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/category/request"
	"github.com/irvansn/go-find-helpers/controllers/category/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryController struct {
	categoryUseCase entities.CategoryUseCaseInterface
}

func (cc *CategoryController) Create(c echo.Context) error {
	var categoryCreate request.CategoryCreate
	if err := c.Bind(&categoryCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	category, errUseCase := cc.categoryUseCase.Create(categoryCreate.CreateToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	categoryResponse := response.FromUseCase(&category)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Category created!", categoryResponse))
}

func NewCategoryController(categoryUseCase entities.CategoryUseCaseInterface) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUseCase,
	}
}
