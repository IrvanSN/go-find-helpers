package auth

import "github.com/google/uuid"

type Auth struct {
	ID           uuid.UUID `gorm:"type:varchar(100);"`
	UserID       uuid.UUID `gorm:"type:varchar(100);"`
	Email        string    `gorm:"type:varchar(255);not null;unique"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
}
