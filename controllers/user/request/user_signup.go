package request

import "github.com/irvansn/go-find-helpers/entities"

type UserSignUp struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

func (r *UserSignUp) ToEntities() *entities.User {
	return &entities.User{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		PhoneNumber: r.PhoneNumber,
		Auth: entities.Auth{
			Email:        r.Email,
			PasswordHash: r.Password,
		},
		Role: r.Role,
	}
}
