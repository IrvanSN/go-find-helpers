package address

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/controllers/address/request"
	"github.com/irvansn/go-find-helpers/controllers/address/response"
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AddressController struct {
	addressUseCase entities.AddressUseCaseInterface
}

func (ac *AddressController) Update(c echo.Context) error {
	var addressUpdateRequest request.AddressDetailRequest
	if err := c.Bind(&addressUpdateRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	addressId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	addressUpdateRequest.ID = addressId

	address, errUseCase := ac.addressUseCase.Update(addressUpdateRequest.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	categoryResponse := response.FromUseCase(&address)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success update Address data!", categoryResponse))
}

func (ac *AddressController) Delete(c echo.Context) error {
	var addressDeleteRequest entities.Address

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	addressId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	addressDeleteRequest.ID = addressId

	address, errUseCase := ac.addressUseCase.Delete(&addressDeleteRequest, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	categoryResponse := response.FromUseCase(&address)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success deleted Address data!", categoryResponse))
}

func NewAddressController(addressUseCase entities.AddressUseCaseInterface) *AddressController {
	return &AddressController{
		addressUseCase: addressUseCase,
	}
}
