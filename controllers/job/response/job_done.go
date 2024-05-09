package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type JobDonePayment struct {
	PaymentID uuid.UUID `json:"payment_id"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
}

type JobDoneTransaction struct {
	TransactionID uuid.UUID      `json:"transaction_id"`
	Type          string         `json:"type"`
	UserID        uuid.UUID      `json:"user_id"`
	SubTotal      float64        `json:"sub_total"`
	Tax           float64        `json:"tax"`
	Total         float64        `json:"total"`
	Payment       JobDonePayment `json:"payment"`
}

type JobMarkDoneResponse struct {
	JobID        uuid.UUID            `json:"job_id"`
	Title        string               `json:"title"`
	Description  string               `json:"description"`
	Status       string               `json:"status"`
	Transactions []JobDoneTransaction `json:"transactions"`
}

func MarkDoneResponseFromUseCase(job *entities.Job) *JobMarkDoneResponse {
	transactions := make([]JobDoneTransaction, len(job.Transactions))
	for i, transaction := range job.Transactions {
		transactions[i] = JobDoneTransaction{
			TransactionID: transaction.ID,
			Type:          transaction.Type,
			UserID:        transaction.UserID,
			SubTotal:      transaction.SubTotal,
			Tax:           transaction.Tax,
			Total:         transaction.Total,
			Payment: JobDonePayment{
				PaymentID: transaction.Payment.ID,
				Status:    transaction.Payment.Status,
				Amount:    float64(transaction.Payment.Amount),
			},
		}
	}
	return &JobMarkDoneResponse{
		JobID:       job.ID,
		Title:       job.Title,
		Description: job.Description,
		Status:      job.Status,
	}
}
