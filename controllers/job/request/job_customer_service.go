package request

import (
	"github.com/irvansn/go-find-helpers/entities"
)

type JobCustomerServiceRequest struct {
	Question string `json:"question"`
}

func (r *JobCustomerServiceRequest) JobIdToEntities() *entities.JobCustomerService {
	return &entities.JobCustomerService{
		Question: r.Question,
	}
}
