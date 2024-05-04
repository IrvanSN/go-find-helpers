package category

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type Category struct {
	ID   uuid.UUID `gorm:"type:varchar(100);"`
	Name string    `gorm:"type:varchar(100);not null"`
}

func FromUseCase(category *entities.Category) *Category {
	return &Category{
		ID:   category.ID,
		Name: category.Name,
	}
}

func (category *Category) ToUseCase() *entities.Category {
	return &entities.Category{
		ID:   category.ID,
		Name: category.Name,
	}
}
