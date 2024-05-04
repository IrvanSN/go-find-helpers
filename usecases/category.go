package usecases

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
)

type CategoryUseCase struct {
	repository entities.CategoryRepositoryInterface
}

func NewCategoryUseCase(repository entities.CategoryRepositoryInterface) *CategoryUseCase {
	return &CategoryUseCase{repository: repository}
}

func (c *CategoryUseCase) Create(category *entities.Category) (entities.Category, error) {
	if category.Name == "" {
		return entities.Category{}, constant.ErrEmptyInput
	}

	category.ID = uuid.New()

	if err := c.repository.Create(category); err != nil {
		return entities.Category{}, err
	}

	return *category, nil
}
