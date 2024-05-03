package reward

import (
	"github.com/google/uuid"
	"time"
)

type Reward struct {
	ID                uuid.UUID `gorm:"type:varchar(100);"`
	Type              string    `gorm:"type:varchar(100);not null"`
	Status            string    `gorm:"type:varchar(50);not null"`
	UserID            uuid.UUID `gorm:"type:varchar(100);not null"`
	JobID             uuid.UUID `gorm:"type:varchar(100);not null"`
	SubTotal          float64   `gorm:"type:decimal;not null"`
	Tax               float64   `gorm:"type:decimal;not null"`
	Total             float64   `gorm:"type:decimal;not null"`
	PaymentExternalId string    `gorm:"type:varchar(100);"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}
