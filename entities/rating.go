package entities

import (
	"github.com/google/uuid"
	"time"
)

type Rating struct {
	ID        uuid.UUID
	Star      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
