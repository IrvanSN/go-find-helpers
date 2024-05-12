package thumbnail

import (
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewThumbnailRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) GetPreSignedURL(thumbnail *entities.Thumbnail) error {
	thumbnailDb := FromUseCase(thumbnail)

	if err := thumbnailDb.GetPreSignedURL(); err != nil {
		return err
	}

	*thumbnail = *thumbnailDb.ToUseCase()
	return nil
}
