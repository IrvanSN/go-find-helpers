package entities

import "github.com/google/uuid"

type Payment struct {
	ID            uuid.UUID
	TransactionID uuid.UUID
	ExternalID    string
	Status        string
	Amount        int64
	InvoiceURL    string
}
