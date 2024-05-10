package entities

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/middlewares"
	"time"
)

type Transaction struct {
	ID        uuid.UUID
	Type      string
	UserID    uuid.UUID
	JobID     uuid.UUID
	SubTotal  float64
	Tax       float64
	Total     float64
	Payment   Payment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionRepositoryInterface interface {
	GetAllTransaction(transactions *[]Transaction, transactionType string) error
}

type TransactionUseCaseInterface interface {
	GetAllTransactions(transactions *[]Transaction, transactionType string, user *middlewares.Claims) ([]Transaction, error)
}
