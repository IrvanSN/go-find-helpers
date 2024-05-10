package transaction

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/payment"
	"github.com/irvansn/go-find-helpers/entities"
	"time"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:varchar(100);"`
	Type      string    `gorm:"type:varchar(100);not null"`
	UserID    uuid.UUID `gorm:"type:varchar(100);not null"`
	JobID     uuid.UUID `gorm:"type:varchar(100);not null"`
	SubTotal  float64   `gorm:"type:decimal;not null"`
	Tax       float64   `gorm:"type:decimal;not null"`
	Total     float64   `gorm:"type:decimal;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Payment   payment.Payment
}

func FromUseCase(transaction *entities.Transaction) *Transaction {
	return &Transaction{
		ID:        transaction.ID,
		Type:      transaction.Type,
		UserID:    transaction.UserID,
		JobID:     transaction.JobID,
		SubTotal:  transaction.SubTotal,
		Tax:       transaction.Tax,
		Total:     transaction.Total,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
		Payment: payment.Payment{
			ID:         transaction.Payment.ID,
			Amount:     transaction.Payment.Amount,
			Status:     transaction.Payment.Status,
			InvoiceURL: transaction.Payment.InvoiceURL,
		},
	}
}

func (t *Transaction) ToUseCase() *entities.Transaction {
	return &entities.Transaction{
		ID:        t.ID,
		Type:      t.Type,
		UserID:    t.UserID,
		JobID:     t.JobID,
		SubTotal:  t.SubTotal,
		Tax:       t.Tax,
		Total:     t.Total,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Payment: entities.Payment{
			ID:         t.Payment.ID,
			Amount:     t.Payment.Amount,
			Status:     t.Payment.Status,
			InvoiceURL: t.Payment.InvoiceURL,
		},
	}
}
