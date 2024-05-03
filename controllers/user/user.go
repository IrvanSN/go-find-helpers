package controllers

import (
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/user/request"
	"github.com/irvansn/go-find-helpers/controllers/user/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

type UserController struct {
	userUseCase entities.UseCaseInterface
}

func (uc *UserController) SignUp(c echo.Context) error {
	var userSignUp request.UserSignUp
	errBindData := c.Bind(&userSignUp)
	if errBindData != nil {
		return c.JSON(utils.ConvertResponseCode(errBindData), base.NewErrorResponse(errBindData.Error()))
	}

	user, errUseCase := uc.userUseCase.SignUp(userSignUp.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	token, errTokenCreation := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
		"email":      user.Auth.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"role":       user.Role,
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if errTokenCreation != nil {
		return c.JSON(utils.ConvertResponseCode(errTokenCreation), base.NewErrorResponse(errTokenCreation.Error()))
	}

	userResponse := response.FromUseCase(&user, token)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", userResponse))
}

func NewUserController(userUseCase entities.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}
