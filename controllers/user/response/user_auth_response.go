package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type UserAuthResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	UserToken string    `json:"user_token"`
}

func AuthResponseFromUseCase(user *entities.User, token string) *UserAuthResponse {
	return &UserAuthResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Auth.Email,
		Role:      user.Role,
		UserToken: token,
	}
}
