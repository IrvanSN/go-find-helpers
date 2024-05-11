package address

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
	"time"
)

type Address struct {
	ID        uuid.UUID `gorm:"type:varchar(100);"`
	Address   string    `gorm:"type:varchar(255);not null"`
	City      string    `gorm:"type:varchar(100);not null"`
	State     string    `gorm:"type:varchar(50);not null"`
	ZipCode   string    `gorm:"type:varchar(10);not null"`
	Country   string    `gorm:"type:varchar(50);not null"`
	Longitude string    `gorm:"type:varchar(100);not null"`
	Latitude  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

func FromUseCase(address *entities.Address) *Address {
	return &Address{
		ID:        address.ID,
		Address:   address.Address,
		City:      address.City,
		State:     address.State,
		ZipCode:   address.ZipCode,
		Country:   address.Country,
		Longitude: address.Longitude,
		Latitude:  address.Latitude,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}
}

func (address *Address) ToUseCase() *entities.Address {
	return &entities.Address{
		ID:        address.ID,
		Address:   address.Address,
		City:      address.City,
		State:     address.State,
		ZipCode:   address.ZipCode,
		Country:   address.Country,
		Longitude: address.Longitude,
		Latitude:  address.Latitude,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}
}
