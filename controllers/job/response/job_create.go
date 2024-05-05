package response

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type JobCreateResponse struct {
	JobId      uuid.UUID `json:"job_id"`
	Amount     int64     `json:"amount"`
	PaymentId  uuid.UUID `json:"payment_id"`
	PaymentUrl string    `json:"payment_url"`
}

func CreateResponseFromUseCase(job *entities.Job) *JobCreateResponse {
	return &JobCreateResponse{
		JobId:      job.ID,
		PaymentId:  job.Transactions[0].Payment.ID,
		PaymentUrl: job.Transactions[0].Payment.InvoiceURL,
		Amount:     job.Transactions[0].Payment.Amount,
	}
}
