package rating

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRatingRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r Repo) Create(rating *entities.Rating) error {
	ratingDb := FromUseCase(rating)

	if err := r.DB.Create(&ratingDb).Error; err != nil {
		return err
	}

	*rating = *ratingDb.ToUseCase()
	return nil
}

func (r Repo) Update(rating *entities.Rating) error {
	ratingDb := FromUseCase(rating)

	db := r.DB.Where("id = ?", ratingDb.ID).Updates(&ratingDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*rating = *ratingDb.ToUseCase()
	return nil
}

func (r Repo) Delete(rating *entities.Rating) error {
	ratingDb := FromUseCase(rating)

	db := r.DB.Delete(&ratingDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*rating = *ratingDb.ToUseCase()
	return nil
}

func (r Repo) GetAll(ratings *[]entities.Rating, ratingUserId uuid.UUID) error {
	var ratingDb []Rating

	if err := r.DB.Where("to_user_id = ?", ratingUserId).Find(&ratingDb).Error; err != nil {
		return err
	}

	for _, _rating := range ratingDb {
		*ratings = append(*ratings, *_rating.ToUseCase())
	}
	return nil
}
