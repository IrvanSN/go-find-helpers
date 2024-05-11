package user

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/address"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/auth"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/payment"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/transaction"
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
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
	Role           string            `gorm:"type:varchar(100);not null"`
	CreatedAt      time.Time         `gorm:"autoCreateTime"`
	UpdatedAt      time.Time         `gorm:"autoUpdateTime"`
	Addresses      []address.Address `gorm:"many2many:user_addresses;"`
	Transactions   []transaction.Transaction
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(user *entities.User) *User {
	addresses := make([]address.Address, len(user.Addresses))
	for i, _address := range user.Addresses {
		addresses[i] = address.Address{
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

	transactions := make([]transaction.Transaction, len(user.Transactions))
	for i, _transaction := range user.Transactions {
		transactions[i] = transaction.Transaction{
			ID:       _transaction.ID,
			Type:     _transaction.Type,
			UserID:   _transaction.UserID,
			JobID:    _transaction.JobID,
			SubTotal: _transaction.SubTotal,
			Tax:      _transaction.Tax,
			Total:    _transaction.Total,
			Payment: payment.Payment{
				ID:         _transaction.Payment.ID,
				Amount:     _transaction.Payment.Amount,
				Status:     _transaction.Payment.Status,
				InvoiceURL: _transaction.Payment.InvoiceURL,
			},
			CreatedAt: _transaction.CreatedAt,
			UpdatedAt: _transaction.UpdatedAt,
		}
	}

	return &User{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		CurrentRating:  user.CurrentRating,
		CurrentBalance: user.CurrentBalance,
		Auth: auth.Auth{
			ID:           user.Auth.ID,
			UserID:       user.ID,
			Email:        user.Auth.Email,
			PasswordHash: user.Auth.PasswordHash,
		},
		Transactions: transactions,
		Role:         user.Role,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Addresses:    addresses,
	}
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

	transactions := make([]entities.Transaction, len(u.Transactions))
	for i, _transaction := range u.Transactions {
		transactions[i] = entities.Transaction{
			ID:       _transaction.ID,
			Type:     _transaction.Type,
			UserID:   _transaction.UserID,
			JobID:    _transaction.JobID,
			SubTotal: _transaction.SubTotal,
			Tax:      _transaction.Tax,
			Total:    _transaction.Total,
			Payment: entities.Payment{
				ID:         _transaction.Payment.ID,
				Amount:     _transaction.Payment.Amount,
				Status:     _transaction.Payment.Status,
				InvoiceURL: _transaction.Payment.InvoiceURL,
			},
			CreatedAt: _transaction.CreatedAt,
			UpdatedAt: _transaction.UpdatedAt,
		}
	}
	return &entities.User{
		ID:             u.ID,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		PhoneNumber:    u.PhoneNumber,
		CurrentRating:  u.CurrentRating,
		CurrentBalance: u.CurrentBalance,
		Role:           u.Role,
		Auth: entities.Auth{
			ID:           u.Auth.ID,
			Email:        u.Auth.Email,
			PasswordHash: u.Auth.PasswordHash,
		},
		Transactions: transactions,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		Addresses:    addresses,
	}
}
