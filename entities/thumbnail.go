package entities

import "github.com/google/uuid"

type Thumbnail struct {
	ID           uuid.UUID
	JobID        uuid.UUID
	ImageKey     string
	Description  string
	PreSignedURL string
}

type ThumbnailRepositoryInterface interface {
	GetPreSignedURL(thumbnail *Thumbnail) error
}

type ThumbnailUseCaseInterface interface {
	GetPreSignedURL(thumbnail *Thumbnail) (Thumbnail, error)
}
