package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type AddressResponse struct {
	AddressId uuid.UUID `json:"address_id"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
}

type CategoryResponse struct {
	CategoryId uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
}

type ThumbnailResponse struct {
	ThumbnailId uuid.UUID `json:"thumbnail_id"`
	ImageKey    string    `json:"image_key"`
	Description string    `json:"description"`
}

type JobResponse struct {
	JobId          uuid.UUID           `json:"job_id"`
	Title          string              `json:"title"`
	Description    string              `json:"description"`
	RewardEarned   float64             `json:"reward_earned"`
	FromAddress    AddressResponse     `json:"from_address"`
	ToAddress      AddressResponse     `json:"to_address"`
	Status         string              `json:"status"`
	HelperRequired uint                `json:"helper_required"`
	Category       CategoryResponse    `json:"category"`
	Thumbnails     []ThumbnailResponse `json:"thumbnails"`
}

type JobGetAllResponse struct {
	Jobs []JobResponse `json:"jobs"`
}

func GetAllResponseFromUseCase(jobs *[]entities.Job) *JobGetAllResponse {
	var jobResponses []JobResponse

	for _, job := range *jobs {
		fromAddress := AddressResponse{
			AddressId: job.FromAddress.ID,
			Address:   job.FromAddress.Address,
			City:      job.FromAddress.City,
			State:     job.FromAddress.State,
			ZipCode:   job.FromAddress.ZipCode,
			Country:   job.FromAddress.Country,
			Longitude: job.FromAddress.Longitude,
			Latitude:  job.FromAddress.Latitude,
		}

		toAddress := AddressResponse{
			AddressId: job.ToAddress.ID,
			Address:   job.ToAddress.Address,
			City:      job.ToAddress.City,
			State:     job.ToAddress.State,
			ZipCode:   job.ToAddress.ZipCode,
			Country:   job.ToAddress.Country,
			Longitude: job.ToAddress.Longitude,
			Latitude:  job.ToAddress.Latitude,
		}

		category := CategoryResponse{
			CategoryId: job.Category.ID,
			Name:       job.Category.Name,
		}

		var thumbnails []ThumbnailResponse
		for _, thumbnail := range job.Thumbnails {
			thumbnails = append(thumbnails, ThumbnailResponse{
				ThumbnailId: thumbnail.ID,
				ImageKey:    thumbnail.ImageKey,
				Description: thumbnail.Description,
			})
		}

		jobResponse := JobResponse{
			JobId:          job.ID,
			Title:          job.Title,
			Description:    job.Description,
			RewardEarned:   job.RewardEarned,
			FromAddress:    fromAddress,
			ToAddress:      toAddress,
			Status:         job.Status,
			HelperRequired: job.HelperRequired,
			Category:       category,
			Thumbnails:     thumbnails,
		}

		jobResponses = append(jobResponses, jobResponse)
	}

	return &JobGetAllResponse{
		Jobs: jobResponses,
	}
}
