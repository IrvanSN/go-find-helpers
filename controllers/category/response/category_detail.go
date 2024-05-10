package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type CategoryDetailResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func FromUseCase(category *entities.Category) *CategoryDetailResponse {
	return &CategoryDetailResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
