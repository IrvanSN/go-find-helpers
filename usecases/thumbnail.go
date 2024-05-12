package usecases

import "github.com/irvansn/go-find-helpers/entities"

type ThumbnailUseCase struct {
	repository entities.ThumbnailRepositoryInterface
}

func NewThumbnailUseCase(repository entities.ThumbnailRepositoryInterface) *ThumbnailUseCase {
	return &ThumbnailUseCase{repository: repository}
}

func (c *ThumbnailUseCase) GetPreSignedURL(thumbnail *entities.Thumbnail) (entities.Thumbnail, error) {
	if err := c.repository.GetPreSignedURL(thumbnail); err != nil {
		return entities.Thumbnail{}, err
	}

	return *thumbnail, nil
}
