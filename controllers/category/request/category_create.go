package request

import "github.com/irvansn/go-find-helpers/entities"

type CategoryCreate struct {
	Name string `json:"name"`
}

func (r *CategoryCreate) CreateToEntities() *entities.Category {
	return &entities.Category{
		Name: r.Name,
	}
}
