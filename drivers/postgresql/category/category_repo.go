package category

import (
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(category *entities.Category) error {
	categoryDb := FromUseCase(category)

	if err := r.DB.Create(&categoryDb).Error; err != nil {
		return err
	}

	*category = *categoryDb.ToUseCase()
	return nil
}
