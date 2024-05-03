package request

import "github.com/irvansn/go-find-helpers/entities"

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (r *UserSignIn) SignInToEntities() *entities.User {
	return &entities.User{
		Auth: entities.Auth{
			Email:        r.Email,
			PasswordHash: r.Password,
		},
		Role: r.Role,
	}
}
