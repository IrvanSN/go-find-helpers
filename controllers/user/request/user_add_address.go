package request

import "github.com/irvansn/go-find-helpers/entities"

type UserAddAddress struct {
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
	Country   string `json:"country"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func (r *UserAddAddress) AddAddressToEntities() *entities.User {
	user := &entities.User{}
	newAddress := entities.Address{
		Address:   r.Address,
		City:      r.City,
		State:     r.State,
		ZipCode:   r.ZipCode,
		Country:   r.Country,
		Longitude: r.Longitude,
		Latitude:  r.Latitude,
	}
	user.Addresses = append(user.Addresses, newAddress)
	return user
}
