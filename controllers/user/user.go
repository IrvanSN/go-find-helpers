package controllers

import (
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/user/request"
	"github.com/irvansn/go-find-helpers/controllers/user/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
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

	tokenExpires := jwt2.NewNumericDate(time.Now().Add(time.Hour * 24))

	claims := &middlewares.Claims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: tokenExpires,
		},
		ID:        user.ID,
		Email:     user.Auth.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}

	token, errTokenCreation := jwt2.NewWithClaims(jwt2.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if errTokenCreation != nil {
		return c.JSON(utils.ConvertResponseCode(errTokenCreation), base.NewErrorResponse(errTokenCreation.Error()))
	}

	userResponse := response.AuthResponseFromUseCase(&user, token)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success sign-up!", userResponse))
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

	tokenExpires := jwt2.NewNumericDate(time.Now().Add(time.Hour * 24))

	claims := &middlewares.Claims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: tokenExpires,
		},
		ID:        user.ID,
		Email:     user.Auth.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}

	token, errTokenCreation := jwt2.NewWithClaims(jwt2.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if errTokenCreation != nil {
		return c.JSON(utils.ConvertResponseCode(errTokenCreation), base.NewErrorResponse(errTokenCreation.Error()))
	}

	userResponse := response.AuthResponseFromUseCase(&user, token)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success sign-in!", userResponse))
}

func (uc *UserController) AddAddress(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var userAddAddress request.UserAddAddress
	if err := c.Bind(&userAddAddress); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	user, errUseCase := uc.userUseCase.AddAddress(userAddAddress.AddAddressToEntities(), userData.ID)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	userResponse := response.AddressResponseFromUseCase(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success add address!", userResponse))
}

func (uc *UserController) GetAllAddresses(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	user, errUseCase := uc.userUseCase.GetAllAddresses(&entities.User{ID: userData.ID})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	userResponse := response.AllAddressesResponseFromUseCase(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all addresses!", userResponse))
}

func NewUserController(userUseCase entities.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}
