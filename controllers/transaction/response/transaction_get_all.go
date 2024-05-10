package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
	"time"
)

type PaymentResponse struct {
	ID     uuid.UUID `json:"payment_id"`
	Status string    `json:"status"`
	Amount float64   `json:"amount"`
}

type TransactionResponse struct {
	ID        uuid.UUID       `json:"id"`
	Type      string          `json:"type"`
	UserID    uuid.UUID       `json:"user_id"`
	JobID     uuid.UUID       `json:"job_id"`
	SubTotal  float64         `json:"sub_total"`
	Tax       float64         `json:"tax"`
	Total     float64         `json:"total"`
	Payment   PaymentResponse `json:"payment"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type TransactionGetAllResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}

func GetAllTransactions(transactions *[]entities.Transaction) TransactionGetAllResponse {
	var respTransactions []TransactionResponse
	for _, transaction := range *transactions {
		respTransaction := TransactionResponse{
			ID:       transaction.ID,
			Type:     transaction.Type,
			UserID:   transaction.UserID,
			JobID:    transaction.JobID,
			SubTotal: transaction.SubTotal,
			Tax:      transaction.Tax,
			Total:    transaction.Total,
			Payment: PaymentResponse{
				ID:     transaction.Payment.ID,
				Status: transaction.Payment.Status,
				Amount: float64(transaction.Payment.Amount),
			},
			CreatedAt: transaction.CreatedAt,
			UpdatedAt: transaction.UpdatedAt,
		}
		respTransactions = append(respTransactions, respTransaction)
	}

	return TransactionGetAllResponse{
		Transactions: respTransactions,
	}
}
