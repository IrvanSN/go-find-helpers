package controllers

import (
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/user/request"
	"github.com/irvansn/go-find-helpers/controllers/user/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userUseCase entities.UseCaseInterface
}

func (uc *UserController) SignUp(c echo.Context) error {
	var userSignUp request.UserSignUp
	err := c.Bind(&userSignUp)
	if err != nil {
		return err
	}

	user, err := uc.userUseCase.SignUp(userSignUp.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.FromUseCase(&user, "")
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", userResponse))
}

func NewUserController(userUseCase entities.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}
