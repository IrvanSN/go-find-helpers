package controllers

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/category/request"
	"github.com/irvansn/go-find-helpers/controllers/category/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryController struct {
	categoryUseCase entities.CategoryUseCaseInterface
}

func (cc *CategoryController) Create(c echo.Context) error {
	var categoryCreate request.CategoryDetailRequest
	if err := c.Bind(&categoryCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	category, errUseCase := cc.categoryUseCase.Create(categoryCreate.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	categoryResponse := response.FromUseCase(&category)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Category created!", categoryResponse))
}

func (cc *CategoryController) Update(c echo.Context) error {
	var categoryUpdateRequest request.CategoryDetailRequest
	if err := c.Bind(&categoryUpdateRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	categoryId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	categoryUpdateRequest.ID = categoryId

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	category, errUseCase := cc.categoryUseCase.Update(categoryUpdateRequest.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	categoryResponse := response.FromUseCase(&category)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Category updated!", categoryResponse))
}

func (cc *CategoryController) Delete(c echo.Context) error {
	var categoryDeleteRequest entities.Category
	if err := c.Bind(&categoryDeleteRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	categoryId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	categoryDeleteRequest.ID = categoryId

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	category, errUseCase := cc.categoryUseCase.Delete(&categoryDeleteRequest, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	categoryResponse := response.FromUseCase(&category)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Category deleted!", categoryResponse))
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	categories, errUseCase := cc.categoryUseCase.GetAll(&[]entities.Category{})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	categoryGetAllResponse := response.SliceFromUseCase(&categories)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Berhasil mendapatkan semua kategori yang tersedia!", categoryGetAllResponse))
}

func NewCategoryController(categoryUseCase entities.CategoryUseCaseInterface) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUseCase,
	}
}
