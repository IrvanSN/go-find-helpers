package reward

import (
	"github.com/google/uuid"
	"time"
)

type Reward struct {
	ID                uuid.UUID `gorm:"type:varchar(100);" json:"id"`
	Type              string    `gorm:"type:varchar(100);not null" json:"type"`
	Status            string    `gorm:"type:varchar(50);not null" json:"status"`
	UserID            uuid.UUID `gorm:"type:varchar(100);not null" json:"user_id"`
	JobID             uuid.UUID `gorm:"type:varchar(100);not null" json:"job_id"`
	SubTotal          float64   `gorm:"type:decimal;not null" json:"sub_total"`
	Tax               float64   `gorm:"type:decimal;not null" json:"tax"`
	Total             float64   `gorm:"type:decimal;not null" json:"total"`
	PaymentExternalId string    `gorm:"type:varchar(100);" json:"payment_external_id"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
