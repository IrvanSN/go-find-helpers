package entities

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID        uuid.UUID
	Type      string
	UserID    uuid.UUID
	JobID     uuid.UUID
	SubTotal  float64
	Tax       float64
	Total     float64
	Payment   Payment
	CreatedAt time.Time
	UpdatedAt time.Time
}
