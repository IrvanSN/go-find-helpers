package entities

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/middlewares"
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
	Rewards        []Transaction
	Ratings        []Rating
	Addresses      []Address
}

type UserRepositoryInterface interface {
	SignUp(user *User) error
	SignIn(user *User) error
	AddAddress(user *User) error
	GetAllAddresses(user *User) error
	Find(user *User) error
	Update(user *User) error
}

type UserUseCaseInterface interface {
	SignUp(user *User) (User, error)
	SignIn(user *User) (User, error)
	AddAddress(user *User, userId uuid.UUID) (User, error)
	Find(user *User) (User, error)
	GetAllAddresses(user *User) (User, error)
	Update(user *User, userRequest *middlewares.Claims) (User, error)
}
