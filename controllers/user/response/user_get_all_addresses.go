package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type userAddress struct {
	ID        uuid.UUID `json:"id"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
}

type GetAllAddressesResponse struct {
	UserID    uuid.UUID     `json:"user_id"`
	Addresses []userAddress `json:"addresses"`
}

func AllAddressesResponseFromUseCase(user *entities.User) *GetAllAddressesResponse {
	addresses := make([]userAddress, len(user.Addresses))
	for i, address := range user.Addresses {
		addresses[i] = userAddress{
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
	return &GetAllAddressesResponse{
		UserID:    user.ID,
		Addresses: addresses,
	}
}
