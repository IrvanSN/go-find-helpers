package usecases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"net/http"
	"os"
)

type JobUseCase struct {
	repository entities.JobRepositoryInterface
}

func NewJobUseCase(repository entities.JobRepositoryInterface) *JobUseCase {
	return &JobUseCase{
		repository: repository,
	}
}

type xdtInvoicePayload struct {
	ExternalID string `json:"external_id"`
	Amount     int    `json:"amount"`
}

type xdtInvoiceResponse struct {
	ID         string `json:"id"`
	InvoiceURL string `json:"invoice_url"`
}

func (j *JobUseCase) Create(job *entities.Job, userId uuid.UUID) (entities.Job, error) {
	if job.Title == "" || job.Description == "" {
		return entities.Job{}, constant.ErrEmptyInput
	}

	if job.RewardEarned < 10000 || job.HelperRequired < 1 {
		return entities.Job{}, constant.ErrInvalidRequest
	}

	job.ID = uuid.New()
	job.Status = "CLOSED"
	job.User.ID = userId

	transaction := entities.Transaction{}
	transaction.ID = uuid.New()

	subTotal := job.RewardEarned * float64(job.HelperRequired)
	tax := (subTotal / 100) * 5

	url := "https://api.xendit.co/v2/invoices"
	method := "POST"

	var payload xdtInvoicePayload
	payload.ExternalID = transaction.ID.String()
	payload.Amount = int(subTotal + tax)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return entities.Job{}, constant.ErrInternalServer
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return entities.Job{}, constant.ErrInternalServer
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+os.Getenv("XDT_SECRET_API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return entities.Job{}, constant.ErrInternalServer
	}
	defer res.Body.Close()

	var response xdtInvoiceResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return entities.Job{}, constant.ErrInternalServer
	}

	transaction.Type = "MONEY_OUT"
	transaction.Status = "PENDING"
	transaction.User = entities.User{ID: userId}
	transaction.Job = entities.Job{ID: job.ID}
	transaction.SubTotal = subTotal
	transaction.Tax = tax
	transaction.Total = subTotal + tax
	transaction.PaymentExternalId = response.InvoiceURL

	job.Transactions = append(job.Transactions, transaction)

	if err := j.repository.Create(job); err != nil {
		return entities.Job{}, err
	}

	return *job, nil
}
