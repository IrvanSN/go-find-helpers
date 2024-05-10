package category

import (
	"github.com/irvansn/go-find-helpers/constant"
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

func (r *Repo) Update(category *entities.Category) error {
	categoryDb := FromUseCase(category)

	db := r.DB.Where("id = ?", categoryDb.ID).Updates(&categoryDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*category = *categoryDb.ToUseCase()
	return nil
}

func (r *Repo) Delete(category *entities.Category) error {
	categoryDb := FromUseCase(category)

	db := r.DB.Delete(&categoryDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*category = *categoryDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(categories *[]entities.Category) error {
	var categoryDb []Category

	if err := r.DB.Find(&categoryDb).Error; err != nil {
		return err
	}

	for _, category := range categoryDb {
		*categories = append(*categories, *category.ToUseCase())
	}
	return nil
}
