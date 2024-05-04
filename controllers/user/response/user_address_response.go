package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type UserAddressResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
}

func AddressResponseFromUseCase(u *entities.User) *UserAddressResponse {
	address := &UserAddressResponse{}
	address.ID = u.Addresses[0].ID
	address.UserID = u.ID
	address.Address = u.Addresses[0].Address
	address.City = u.Addresses[0].City
	address.State = u.Addresses[0].State
	address.ZipCode = u.Addresses[0].ZipCode
	address.Country = u.Addresses[0].Country
	address.Longitude = u.Addresses[0].Longitude
	address.Latitude = u.Addresses[0].Latitude
	return address
}
