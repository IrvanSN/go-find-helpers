package thumbnail

import "gorm.io/gorm"

type Repo struct {
	DB *gorm.DB
}

func NewThumbnailRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}
