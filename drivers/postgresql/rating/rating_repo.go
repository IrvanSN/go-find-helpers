package rating

import "gorm.io/gorm"

type Repo struct {
	DB *gorm.DB
}

func NewRatingRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}
