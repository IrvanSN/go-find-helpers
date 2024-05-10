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

type UserGetAllTransactionsResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}

func GetAllTransactions(user *entities.User) UserGetAllTransactionsResponse {
	transactions := make([]TransactionResponse, len(user.Transactions))
	for i, transaction := range user.Transactions {
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

	return UserGetAllTransactionsResponse{
		Transactions: transactions,
	}
}
