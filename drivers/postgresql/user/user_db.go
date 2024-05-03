package user

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/address"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/auth"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/job"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/rating"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/reward"
	"github.com/irvansn/go-find-helpers/entities"
	"time"
)

type User struct {
	ID             uuid.UUID `gorm:"type:varchar(100);" json:"id"`
	FirstName      string    `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName       string    `gorm:"type:varchar(100);not null" json:"last_name"`
	PhoneNumber    string    `gorm:"type:varchar(20);not null" json:"phone_number"`
	CurrentRating  float32   `gorm:"type:decimal;not null" json:"current_rating"`
	CurrentBalance float64   `gorm:"type:decimal;not null" json:"current_balance"`
	Auth           auth.Auth
	Role           string    `gorm:"type:varchar(100);not null" json:"role"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Jobs           []job.Job
	Rewards        []reward.Reward
	Ratings        []rating.Rating
	Addresses      []address.Address `gorm:"many2many:user_addresses;"`
}

func FromUseCase(user *entities.User) *User {
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

func (user *User) ToUseCase() *entities.User {
	return &entities.User{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		CurrentBalance: user.CurrentBalance,
		CurrentRating:  user.CurrentRating,
		Role:           user.Role,
		Auth: entities.Auth{
			ID:           user.Auth.ID,
			Email:        user.Auth.Email,
			PasswordHash: user.Auth.PasswordHash,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
