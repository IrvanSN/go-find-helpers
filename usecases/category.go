package usecases

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
)

type CategoryUseCase struct {
	repository entities.CategoryRepositoryInterface
}

func NewCategoryUseCase(repository entities.CategoryRepositoryInterface) *CategoryUseCase {
	return &CategoryUseCase{repository: repository}
}

func (c *CategoryUseCase) Create(category *entities.Category, user *middlewares.Claims) (entities.Category, error) {
	if user.Role != "ADMIN" {
		return entities.Category{}, constant.ErrNotAuthorized
	}

	if category.Name == "" {
		return entities.Category{}, constant.ErrEmptyInput
	}

	category.ID = uuid.New()

	if err := c.repository.Create(category); err != nil {
		return entities.Category{}, err
	}

	return *category, nil
}

func (c *CategoryUseCase) Update(category *entities.Category, user *middlewares.Claims) (entities.Category, error) {
	if user.Role != "ADMIN" {
		return entities.Category{}, constant.ErrNotAuthorized
	}

	if category.ID == uuid.Nil || category.Name == "" {
		return entities.Category{}, constant.ErrEmptyInput
	}

	if err := c.repository.Update(category); err != nil {
		return entities.Category{}, err
	}

	return *category, nil
}

func (c *CategoryUseCase) Delete(category *entities.Category, user *middlewares.Claims) (entities.Category, error) {
	if user.Role != "ADMIN" {
		return entities.Category{}, constant.ErrNotAuthorized
	}

	if category.ID == uuid.Nil {
		return entities.Category{}, constant.ErrEmptyInput
	}

	if err := c.repository.Delete(category); err != nil {
		return entities.Category{}, err
	}

	return *category, nil
}

func (c *CategoryUseCase) GetAll(category *[]entities.Category) ([]entities.Category, error) {
	if err := c.repository.GetAll(category); err != nil {
		return []entities.Category{}, err
	}

	return *category, nil
}
