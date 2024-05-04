package entities

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID
	Name string
}

type CategoryRepositoryInterface interface {
	Create(category *Category) error
}

type CategoryUseCaseInterface interface {
	Create(category *Category) (Category, error)
}
