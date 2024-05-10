package entities

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/middlewares"
)

type Category struct {
	ID   uuid.UUID
	Name string
}

type CategoryRepositoryInterface interface {
	Create(category *Category) error
	Update(category *Category) error
	Delete(category *Category) error
	GetAll(categories *[]Category) error
}

type CategoryUseCaseInterface interface {
	Create(category *Category, user *middlewares.Claims) (Category, error)
	Update(category *Category, user *middlewares.Claims) (Category, error)
	Delete(category *Category, user *middlewares.Claims) (Category, error)
	GetAll(categories *[]Category) ([]Category, error)
}
