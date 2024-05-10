package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type PaymentResponse struct {
	PaymentId uuid.UUID `json:"payment_id"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
}

type TransactionResponse struct {
	TransactionId uuid.UUID       `json:"transaction_id"`
	Type          string          `json:"type"`
	UserID        uuid.UUID       `json:"user_id"`
	SubTotal      float64         `json:"sub_total"`
	Tax           float64         `json:"tax"`
	Total         float64         `json:"total"`
	Payment       PaymentResponse `json:"payment"`
}

type JobStatusUpdateResponse struct {
	JobId        uuid.UUID             `json:"job_id"`
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	RewardEarned float64               `json:"reward_earned"`
	Status       string                `json:"status"`
	FromAddress  AddressResponse       `json:"from_address"`
	ToAddress    AddressResponse       `json:"to_address"`
	Category     CategoryResponse      `json:"category"`
	Transactions []TransactionResponse `json:"transactions"`
	Thumbnails   []ThumbnailResponse   `json:"thumbnails"`
}

func DetailResponseFromUseCase(job *entities.Job) *JobStatusUpdateResponse {
	transactions := make([]TransactionResponse, len(job.Transactions))
	for i, transaction := range job.Transactions {
		transactions[i] = TransactionResponse{
			TransactionId: transaction.ID,
			Type:          transaction.Type,
			UserID:        transaction.UserID,
			SubTotal:      transaction.SubTotal,
			Tax:           transaction.Tax,
			Total:         transaction.Total,
			Payment: PaymentResponse{
				PaymentId: transaction.Payment.ID,
				Status:    transaction.Payment.Status,
				Amount:    float64(transaction.Payment.Amount),
			},
		}
	}

	thumbnails := make([]ThumbnailResponse, len(job.Thumbnails))
	for i, thumbnail := range job.Thumbnails {
		thumbnails[i] = ThumbnailResponse{
			ThumbnailId: thumbnail.ID,
			ImageKey:    thumbnail.ImageKey,
			Description: thumbnail.Description,
		}
	}

	return &JobStatusUpdateResponse{
		JobId:        job.ID,
		Title:        job.Title,
		Description:  job.Description,
		Status:       job.Status,
		RewardEarned: job.RewardEarned,
		FromAddress: AddressResponse{
			AddressId: job.FromAddress.ID,
			Address:   job.FromAddress.Address,
			City:      job.FromAddress.City,
			State:     job.FromAddress.State,
			ZipCode:   job.FromAddress.ZipCode,
			Country:   job.FromAddress.Country,
			Longitude: job.FromAddress.Longitude,
			Latitude:  job.FromAddress.Latitude,
		},
		ToAddress: AddressResponse{
			AddressId: job.ToAddress.ID,
			Address:   job.ToAddress.Address,
			City:      job.ToAddress.City,
			State:     job.ToAddress.State,
			ZipCode:   job.ToAddress.ZipCode,
			Country:   job.ToAddress.Country,
			Longitude: job.ToAddress.Longitude,
			Latitude:  job.ToAddress.Latitude,
		},
		Category: CategoryResponse{
			CategoryId: job.Category.ID,
			Name:       job.Category.Name,
		},
		Transactions: transactions,
		Thumbnails:   thumbnails,
	}
}
