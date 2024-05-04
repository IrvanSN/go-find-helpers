package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID
	FirstName      string
	LastName       string
	PhoneNumber    string
	CurrentRating  float32
	CurrentBalance float64
	Auth           Auth
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Jobs           []Job
	Rewards        []Reward
	Ratings        []Rating
	Addresses      []Address
}

type RepositoryInterface interface {
	SignUp(user *User) error
	SignIn(user *User) error
}

type UseCaseInterface interface {
	SignUp(user *User) (User, error)
	SignIn(user *User) (User, error)
}
