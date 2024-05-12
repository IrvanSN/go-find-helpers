package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type RatingDetailResponse struct {
	ID         uuid.UUID `json:"id"`
	Star       int       `json:"star"`
	FromUserID uuid.UUID `json:"from_user_id"`
	ToUserID   uuid.UUID `json:"to_user_id"`
	JobID      uuid.UUID `json:"job_id"`
}

func FromUseCase(rating *entities.Rating) *RatingDetailResponse {
	return &RatingDetailResponse{
		ID:         rating.ID,
		Star:       rating.Star,
		FromUserID: rating.FromUserID,
		ToUserID:   rating.ToUserID,
		JobID:      rating.JobID,
	}
}
