package entities

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/middlewares"
	"time"
)

type Address struct {
	ID        uuid.UUID
	Address   string
	City      string
	State     string
	ZipCode   string
	Country   string
	Longitude string
	Latitude  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddressRepositoryInterface interface {
	Update(address *Address) error
	Delete(address *Address) error
	Get(address *Address) error
}

type AddressUseCaseInterface interface {
	Update(category *Address, user *middlewares.Claims) (Address, error)
	Delete(category *Address, user *middlewares.Claims) (Address, error)
}
