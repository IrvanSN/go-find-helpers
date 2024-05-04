package entities

import (
	"github.com/google/uuid"
	"time"
)

type Reward struct {
	ID                uuid.UUID
	Type              string
	Status            string
	User              User
	Job               Job
	SubTotal          float64
	Tax               float64
	Total             float64
	PaymentExternalId string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
