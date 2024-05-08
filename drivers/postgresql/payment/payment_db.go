package payment

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"net/http"
	"os"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:varchar(100);"`
	TransactionID uuid.UUID `gorm:"type:varchar(100);"`
	Status        string    `gorm:"type:varchar(50);"`
	ExternalID    string    `gorm:"type:varchar(100);"`
	Amount        int64     `gorm:"type:decimal;not null"`
	InvoiceURL    string    `gorm:"type:text;"`
}

type xdtInvoiceResponse struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Amount     int64  `json:"amount"`
	InvoiceURL string `json:"invoice_url"`
}

type xdtInvoicePayload struct {
	ExternalID  string `json:"external_id"`
	Amount      int64  `json:"amount"`
	PayerEmail  string `json:"payer_email"`
	Description string `json:"description"`
}

func (p *Payment) Create(email string) error {
	url := "https://api.xendit.co/v2/invoices"
	method := "POST"

	p.ExternalID = p.ID.String()

	var payload xdtInvoicePayload
	payload.ExternalID = p.ID.String()
	payload.Amount = p.Amount
	payload.PayerEmail = email
	payload.Description = "Find Helpers App Invoice"
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+os.Getenv("XDT_SECRET_API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return constant.ErrPaymentGateway
	}

	var response xdtInvoiceResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	p.ExternalID = response.ID
	p.InvoiceURL = response.InvoiceURL

	return nil
}
