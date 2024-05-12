package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type RatingDetailRequest struct {
	ID         uuid.UUID
	Star       int       `json:"star"`
	FromUserID uuid.UUID `json:"from_user_id"`
	ToUserID   uuid.UUID `json:"to_user_id"`
	JobID      uuid.UUID `json:"job_id"`
}

func (r *RatingDetailRequest) ToEntities() *entities.Rating {
	return &entities.Rating{
		ID:         r.ID,
		Star:       r.Star,
		FromUserID: r.FromUserID,
		ToUserID:   r.ToUserID,
		JobID:      r.JobID,
	}
}
