package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type AddressDetailResponse struct {
	ID        uuid.UUID `json:"id"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
}

func FromUseCase(address *entities.Address) *AddressDetailResponse {
	return &AddressDetailResponse{
		ID:        address.ID,
		Address:   address.Address,
		City:      address.City,
		State:     address.State,
		ZipCode:   address.ZipCode,
		Country:   address.Country,
		Longitude: address.Longitude,
		Latitude:  address.Latitude,
	}
}
