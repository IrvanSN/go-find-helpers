package usecases

import "github.com/irvansn/go-find-helpers/entities"

type RatingUseCase struct {
	repository entities.RatingRepositoryInterface
}

func NewRatingUseCase(repository entities.RatingRepositoryInterface) *RatingUseCase {
	return &RatingUseCase{repository: repository}
}
