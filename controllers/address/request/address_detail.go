package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type AddressDetailRequest struct {
	ID        uuid.UUID
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
	Country   string `json:"country"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func (r *AddressDetailRequest) ToEntities() *entities.Address {
	return &entities.Address{
		ID:        r.ID,
		Address:   r.Address,
		City:      r.City,
		State:     r.State,
		ZipCode:   r.ZipCode,
		Country:   r.Country,
		Longitude: r.Longitude,
		Latitude:  r.Latitude,
	}
}
