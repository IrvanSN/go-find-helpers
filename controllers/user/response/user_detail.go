package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type AuthResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type UserDetailResponse struct {
	ID          uuid.UUID    `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	PhoneNumber string       `json:"phone_number"`
	Auth        AuthResponse `json:"auth"`
	Role        string       `json:"roles"`
}

func UserDetailResponseFromUseCase(u *entities.User) *UserDetailResponse {
	return &UserDetailResponse{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		PhoneNumber: u.PhoneNumber,
		Auth: AuthResponse{
			ID:    u.ID,
			Email: u.Auth.Email,
		},
		Role: u.Role,
	}
}
