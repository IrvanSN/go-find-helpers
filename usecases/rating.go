package usecases

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
)

type RatingUseCase struct {
	repository entities.RatingRepositoryInterface
}

func NewRatingUseCase(repository entities.RatingRepositoryInterface) *RatingUseCase {
	return &RatingUseCase{repository: repository}
}

func (r *RatingUseCase) Create(rating *entities.Rating) (entities.Rating, error) {
	if rating.FromUserID == uuid.Nil || rating.ToUserID == uuid.Nil || rating.JobID == uuid.Nil || rating.Star < 1 {
		return entities.Rating{}, constant.ErrEmptyInput
	}

	rating.ID = uuid.New()

	if err := r.repository.Create(rating); err != nil {
		return entities.Rating{}, err
	}

	var userRatings []entities.Rating
	if err := r.repository.GetAll(&userRatings, rating.ToUserID); err != nil {
		return entities.Rating{}, err
	}

	var totalRatingEarned = 0
	var ratingDataLength = len(userRatings)

	for _, userRating := range userRatings {
		totalRatingEarned += userRating.Star
	}

	if err := r.repository.UpdateUserRating(rating.ToUserID, float32(totalRatingEarned)/float32(ratingDataLength)); err != nil {
		return entities.Rating{}, err
	}

	return *rating, nil
}

func (r *RatingUseCase) Update(rating *entities.Rating) (entities.Rating, error) {
	if err := r.repository.Update(rating); err != nil {
		return entities.Rating{}, err
	}

	return *rating, nil
}

func (r *RatingUseCase) Delete(rating *entities.Rating, user *middlewares.Claims) (entities.Rating, error) {
	if user.Role != "ADMIN" {
		return entities.Rating{}, constant.ErrNotAuthorized
	}

	if err := r.repository.Delete(rating); err != nil {
		return entities.Rating{}, err
	}

	return *rating, nil
}

func (r *RatingUseCase) GetAll(rating *[]entities.Rating, ratingUserId uuid.UUID) ([]entities.Rating, error) {
	if err := r.repository.GetAll(rating, ratingUserId); err != nil {
		return []entities.Rating{}, err
	}

	return *rating, nil
}
