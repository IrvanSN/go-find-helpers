package response

import "github.com/irvansn/go-find-helpers/entities"

type RatingGetAll struct {
	Ratings []RatingDetailResponse `json:"ratings"`
}

func SliceFromUseCase(ratings *[]entities.Rating) *RatingGetAll {
	result := make([]RatingDetailResponse, len(*ratings))
	for i, rating := range *ratings {
		result[i] = RatingDetailResponse{
			ID:         rating.ID,
			Star:       rating.Star,
			FromUserID: rating.FromUserID,
			ToUserID:   rating.ToUserID,
			JobID:      rating.JobID,
		}
	}
	return &RatingGetAll{
		Ratings: result,
	}
}
