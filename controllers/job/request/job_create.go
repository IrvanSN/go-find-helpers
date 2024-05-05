package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type thumbnailRequest struct {
	ImageKey    string `json:"image_key"`
	Description string `json:"description"`
}

type JobCreateRequest struct {
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	FromAddressId  uuid.UUID          `json:"from_address_id"`
	ToAddressId    uuid.UUID          `json:"to_address_id"`
	HelperRequired uint               `json:"helper_required"`
	CategoryId     uuid.UUID          `json:"category_id"`
	RewardEarned   float64            `json:"reward_earned"`
	Thumbnails     []thumbnailRequest `json:"thumbnails"`
}

func (r *JobCreateRequest) JobCreateToEntities() *entities.Job {
	thumbnails := make([]entities.Thumbnail, len(r.Thumbnails))
	for i, t := range r.Thumbnails {
		thumbnails[i] = entities.Thumbnail{
			ImageKey:    t.ImageKey,
			Description: t.Description,
		}
	}

	return &entities.Job{
		Title:       r.Title,
		Description: r.Description,
		FromAddress: entities.Address{
			ID: r.FromAddressId,
		},
		ToAddress: entities.Address{
			ID: r.ToAddressId,
		},
		HelperRequired: r.HelperRequired,
		Category: entities.Category{
			ID: r.CategoryId,
		},
		RewardEarned: r.RewardEarned,
		Thumbnails:   thumbnails,
	}
}
