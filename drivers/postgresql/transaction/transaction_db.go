package transaction

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/payment"
	"time"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:varchar(100);"`
	Type      string    `gorm:"type:varchar(100);not null"`
	UserID    uuid.UUID `gorm:"type:varchar(100);not null"`
	JobID     uuid.UUID `gorm:"type:varchar(100);not null"`
	SubTotal  float64   `gorm:"type:decimal;not null"`
	Tax       float64   `gorm:"type:decimal;not null"`
	Total     float64   `gorm:"type:decimal;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Payment   payment.Payment
}
