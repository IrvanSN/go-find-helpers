package entities

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/middlewares"
	"time"
)

type Job struct {
	ID             uuid.UUID
	Title          string
	Description    string
	RewardEarned   float64
	FromAddress    Address
	ToAddress      Address
	Status         string
	HelperRequired uint
	Category       Category
	User           User
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Transactions   []Transaction
	Thumbnails     []Thumbnail
}

type JobRepositoryInterface interface {
	Create(Job *Job, user *middlewares.Claims) error
}

type JobUseCaseInterface interface {
	Create(Job *Job, user *middlewares.Claims) (Job, error)
}
