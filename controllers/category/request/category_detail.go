package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type CategoryDetailRequest struct {
	ID   uuid.UUID
	Name string `json:"name"`
}

func (r *CategoryDetailRequest) ToEntities() *entities.Category {
	return &entities.Category{
		ID:   r.ID,
		Name: r.Name,
	}
}
