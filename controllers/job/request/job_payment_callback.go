package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type JobPaymentCallbackRequest struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Status     string `json:"status"`
}

func (j *JobPaymentCallbackRequest) JobPaymentCallbackToEntities() *entities.Job {
	externalID, _ := uuid.Parse(j.ExternalID)

	var transactions []entities.Transaction
	transactions = append(transactions, entities.Transaction{
		Payment: entities.Payment{
			ID:         externalID,
			ExternalID: j.ID,
			Status:     j.Status,
		},
	})

	return &entities.Job{
		Transactions: transactions,
	}
}
