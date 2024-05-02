package address

import (
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID        uuid.UUID `json:"id"`
	Address   string    `gorm:"type:varchar(255);not null" json:"address"`
	City      string    `gorm:"type:varchar(100);not null" json:"city"`
	State     string    `gorm:"type:varchar(50);not null" json:"state"`
	ZipCode   string    `gorm:"type:varchar(10);not null" json:"zip_code"`
	Country   string    `gorm:"type:varchar(50);not null" json:"country"`
	Long      string    `gorm:"type:varchar(100);not null" json:"long"`
	Lat       string    `gorm:"type:varchar(100);not null" json:"lat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
