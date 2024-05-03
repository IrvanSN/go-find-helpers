package entities

import "github.com/google/uuid"

type Auth struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
}
