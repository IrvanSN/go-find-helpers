package response

import "github.com/irvansn/go-find-helpers/entities"

type UploadURLResponse struct {
	Key          string `json:"key"`
	PreSignedURL string `json:"pre_signed_url"`
}

func FromUseCase(thumbnail *entities.Thumbnail) *UploadURLResponse {
	return &UploadURLResponse{
		Key:          thumbnail.ImageKey,
		PreSignedURL: thumbnail.PreSignedURL,
	}
}
