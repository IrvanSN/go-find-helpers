package entities

import (
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID        uuid.UUID
	Address   string
	City      string
	State     string
	ZipCode   string
	Country   string
	Long      string
	Lat       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
