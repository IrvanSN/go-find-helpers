package rating

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/rating/request"
	"github.com/irvansn/go-find-helpers/controllers/rating/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RatingController struct {
	ratingUseCase entities.RatingUseCaseInterface
}

func (rc *RatingController) Create(c echo.Context) error {
	var ratingCreateRequest request.RatingDetailRequest
	if err := c.Bind(&ratingCreateRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	rating, errUseCase := rc.ratingUseCase.Create(ratingCreateRequest.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	ratingResponse := response.FromUseCase(&rating)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Rating created!", ratingResponse))
}

func (rc *RatingController) GetAll(c echo.Context) error {
	var ratingGetAll []entities.Rating

	ratingUserId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	ratings, errUseCase := rc.ratingUseCase.GetAll(&ratingGetAll, ratingUserId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	ratingResponse := response.SliceFromUseCase(&ratings)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Successfully retrieve all Ratings data!", ratingResponse))
}

func (rc *RatingController) Update(c echo.Context) error {
	var ratingUpdateRequest request.RatingDetailRequest
	if err := c.Bind(&ratingUpdateRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	ratingId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	ratingUpdateRequest.ID = ratingId

	rating, errUseCase := rc.ratingUseCase.Update(ratingUpdateRequest.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	ratingResponse := response.FromUseCase(&rating)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Rating updated!", ratingResponse))
}

func (rc *RatingController) Delete(c echo.Context) error {
	var ratingDeleteRequest request.RatingDetailRequest

	ratingId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	ratingDeleteRequest.ID = ratingId

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	rating, errUseCase := rc.ratingUseCase.Delete(ratingDeleteRequest.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	ratingResponse := response.FromUseCase(&rating)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Successfully get all user ratings!", ratingResponse))
}

func NewRatingController(ratingUseCase entities.RatingUseCaseInterface) *RatingController {
	return &RatingController{
		ratingUseCase: ratingUseCase,
	}
}
