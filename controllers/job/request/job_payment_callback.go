package request

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
	"strings"
)

type JobPaymentCallbackRequest struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Status     string `json:"status"`
}

func (j *JobPaymentCallbackRequest) JobPaymentCallbackToEntities() *entities.Job {
	splitID := strings.Split(j.ExternalID, ":")
	fmt.Println("splitID", splitID)
	paymentId, _ := uuid.Parse(splitID[1])
	JobId, _ := uuid.Parse(splitID[0])

	var transactions []entities.Transaction
	transactions = append(transactions, entities.Transaction{
		Payment: entities.Payment{
			ID:         paymentId,
			ExternalID: j.ID,
			Status:     j.Status,
		},
	})

	return &entities.Job{
		ID:           JobId,
		Transactions: transactions,
	}
}
