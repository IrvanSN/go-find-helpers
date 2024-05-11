package entities

import (
	"github.com/google/uuid"
	"time"
)

type Rating struct {
	ID        uuid.UUID
	Star      int
	UserID    uuid.UUID
	JobID     uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RatingRepositoryInterface interface {
}

type RatingUseCaseInterface interface {
}
