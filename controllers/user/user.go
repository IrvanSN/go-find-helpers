package controllers

import (
	"fmt"
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
	userUseCase entities.UserUseCaseInterface
}

func (uc *UserController) SignUp(c echo.Context) error {
	var userSignUp request.UserSignUp
	if err := c.Bind(&userSignUp); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	user, errUseCase := uc.userUseCase.SignUp(userSignUp.SignUpToEntities())
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

func (uc *UserController) SignIn(c echo.Context) error {
	var userSignIn request.UserSignIn
	if err := c.Bind(&userSignIn); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	user, errUseCase := uc.userUseCase.SignIn(userSignIn.SignInToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	fmt.Println("user controller", user)

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

func NewUserController(userUseCase entities.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}
