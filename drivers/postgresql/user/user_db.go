package user

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/address"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/auth"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/job"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/rating"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/transaction"
	"github.com/irvansn/go-find-helpers/entities"
	"time"
)

type User struct {
	ID             uuid.UUID `gorm:"type:varchar(100);"`
	FirstName      string    `gorm:"type:varchar(100);not null"`
	LastName       string    `gorm:"type:varchar(100);not null"`
	PhoneNumber    string    `gorm:"type:varchar(20);not null"`
	CurrentRating  float32   `gorm:"type:decimal;not null"`
	CurrentBalance float64   `gorm:"type:decimal;not null"`
	Auth           auth.Auth
	Role           string    `gorm:"type:varchar(100);not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	Jobs           []job.Job
	Rewards        []transaction.Transaction
	Ratings        []rating.Rating
	Addresses      []address.Address `gorm:"many2many:user_addresses;"`
}

func FromUseCase(user *entities.User) *User {
	return &User{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		CurrentRating:  user.CurrentRating,
		CurrentBalance: user.CurrentBalance,
	}
}

func AuthFromUseCase(user *entities.User) *User {
	return &User{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		CurrentBalance: user.CurrentBalance,
		CurrentRating:  user.CurrentRating,
		Role:           user.Role,
		Auth: auth.Auth{
			ID:           user.Auth.ID,
			Email:        user.Auth.Email,
			PasswordHash: user.Auth.PasswordHash,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func AddressFromUseCase(user *entities.User) *User {
	newUser := &User{
		ID: user.ID,
	}
	newAddress := address.Address{
		ID:        user.Addresses[0].ID,
		Address:   user.Addresses[0].Address,
		City:      user.Addresses[0].City,
		State:     user.Addresses[0].State,
		ZipCode:   user.Addresses[0].ZipCode,
		Country:   user.Addresses[0].Country,
		Longitude: user.Addresses[0].Longitude,
		Latitude:  user.Addresses[0].Latitude,
	}
	newUser.Addresses = append(newUser.Addresses, newAddress)
	return newUser
}

func (u *User) ToUseCase() *entities.User {
	addresses := make([]entities.Address, len(u.Addresses))
	for i, _address := range u.Addresses {
		addresses[i] = entities.Address{
			ID:        _address.ID,
			Address:   _address.Address,
			City:      _address.City,
			State:     _address.State,
			ZipCode:   _address.ZipCode,
			Country:   _address.Country,
			Longitude: _address.Longitude,
			Latitude:  _address.Latitude,
			CreatedAt: _address.CreatedAt,
			UpdatedAt: _address.UpdatedAt,
		}
	}
	return &entities.User{
		ID:             u.ID,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		PhoneNumber:    u.PhoneNumber,
		CurrentRating:  u.CurrentRating,
		CurrentBalance: u.CurrentBalance,
		Addresses:      addresses,
	}
}

func (u *User) AuthToUseCase() *entities.User {
	return &entities.User{
		ID:             u.ID,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		PhoneNumber:    u.PhoneNumber,
		CurrentBalance: u.CurrentBalance,
		CurrentRating:  u.CurrentRating,
		Role:           u.Role,
		Auth: entities.Auth{
			ID:           u.Auth.ID,
			Email:        u.Auth.Email,
			PasswordHash: u.Auth.PasswordHash,
		},
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) AddressToUseCase() *entities.User {
	newUser := &entities.User{
		ID: u.ID,
	}
	newAddress := entities.Address{
		ID:        u.Addresses[0].ID,
		Address:   u.Addresses[0].Address,
		City:      u.Addresses[0].City,
		State:     u.Addresses[0].State,
		ZipCode:   u.Addresses[0].ZipCode,
		Country:   u.Addresses[0].Country,
		Longitude: u.Addresses[0].Longitude,
		Latitude:  u.Addresses[0].Latitude,
	}
	newUser.Addresses = append(newUser.Addresses, newAddress)
	return newUser
}
