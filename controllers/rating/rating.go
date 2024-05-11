package rating

import (
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/labstack/echo/v4"
)

type RatingController struct {
	ratingUseCase entities.RatingUseCaseInterface
}

func (rc *RatingController) Create(c echo.Context) error {
	return nil
}

func (rc *RatingController) GetAll(c echo.Context) error {
	return nil
}

func (rc *RatingController) Update(c echo.Context) error {
	return nil
}

func (rc *RatingController) Delete(c echo.Context) error {
	return nil
}

func NewRatingController(ratingUseCase entities.RatingUseCaseInterface) *RatingController {
	return &RatingController{
		ratingUseCase: ratingUseCase,
	}
}
