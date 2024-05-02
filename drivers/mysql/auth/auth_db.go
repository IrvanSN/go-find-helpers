package auth

import "github.com/google/uuid"

type Auth struct {
	ID           uuid.UUID `gorm:"type:varchar(100);" json:"id"`
	UserID       uuid.UUID `gorm:"type:varchar(100);"`
	Email        string    `gorm:"type:varchar(255);unique_index;not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Strategy     string    `gorm:"type:varchar(50);not null" json:"strategy"`
	Provider     string    `gorm:"type:varchar(50);not null" json:"provider"`
}
