package response

import "github.com/irvansn/go-find-helpers/entities"

type UserGetAllResponse struct {
	Users []UserDetailResponse `json:"users"`
}

func SliceFromUseCase(users *[]entities.User) *UserGetAllResponse {
	allUsers := make([]UserDetailResponse, len(*users))
	for i, _user := range *users {
		allUsers[i] = UserDetailResponse{
			ID:          _user.ID,
			FirstName:   _user.FirstName,
			LastName:    _user.LastName,
			PhoneNumber: _user.PhoneNumber,
			Auth: AuthResponse{
				ID:    _user.Auth.ID,
				Email: _user.Auth.Email,
			},
			Role: _user.Role,
		}
	}

	return &UserGetAllResponse{
		Users: allUsers,
	}
}
