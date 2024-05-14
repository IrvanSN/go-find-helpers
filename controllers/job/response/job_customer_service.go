package response

import "github.com/irvansn/go-find-helpers/entities"

type JobCustomerService struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func JobCustomerServiceFromUseCase(job *entities.JobCustomerService) *JobCustomerService {
	return &JobCustomerService{
		Question: job.Question,
		Answer:   job.Answer,
	}
}
