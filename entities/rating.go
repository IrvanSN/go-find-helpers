package entities

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/middlewares"
	"time"
)

type Rating struct {
	ID         uuid.UUID
	Star       int
	FromUserID uuid.UUID
	ToUserID   uuid.UUID
	JobID      uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type RatingRepositoryInterface interface {
	Create(rating *Rating) error
	Update(rating *Rating) error
	Delete(rating *Rating) error
	GetAll(ratings *[]Rating, ratingUserId uuid.UUID) error
}

type RatingUseCaseInterface interface {
	Create(rating *Rating) (Rating, error)
	Update(rating *Rating) (Rating, error)
	Delete(rating *Rating, user *middlewares.Claims) (Rating, error)
	GetAll(ratings *[]Rating, ratingUserId uuid.UUID) ([]Rating, error)
}
