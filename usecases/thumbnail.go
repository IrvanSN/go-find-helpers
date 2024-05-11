package usecases

import "github.com/irvansn/go-find-helpers/entities"

type ThumbnailUseCase struct {
	repository entities.ThumbnailRepositoryInterface
}

func NewThumbnailUseCase(repository entities.ThumbnailRepositoryInterface) *ThumbnailUseCase {
	return &ThumbnailUseCase{repository: repository}
}
