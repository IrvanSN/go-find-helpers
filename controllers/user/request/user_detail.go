package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type UserDetailRequest struct {
	ID          uuid.UUID
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

func (r *UserDetailRequest) ToEntities() *entities.User {
	return &entities.User{
		ID:          r.ID,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Auth:        entities.Auth{Email: r.Email},
		PhoneNumber: r.PhoneNumber,
		Role:        r.Role,
	}
}
