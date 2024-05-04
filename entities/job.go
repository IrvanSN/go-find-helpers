package entities

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	ID             uuid.UUID
	Title          string
	Description    string
	Reward         float64
	FromAddress    Address
	ToAddress      Address
	Status         string
	HelperRequired uint
	Category       Category
	UserID         User
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Rewards        []Reward
	Thumbnails     []Thumbnail
}

type JobRepositoryInterface interface {
	Create(Job *Job) error
}

type JobUseCaseInterface interface {
	Create(Job *Job) (Job, error)
}
