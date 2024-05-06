package response

import "github.com/irvansn/go-find-helpers/entities"

type JobTakeResponse struct {
	JobId       string `json:"job_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Reward      int64  `json:"reward"`
}

func TakeResponseFromUseCase(job *entities.Job) *JobTakeResponse {
	var rewardTotal int64
	for _, transaction := range job.Transactions {
		if transaction.User.ID == job.User.ID && transaction.Type == "MONEY_IN" {
			rewardTotal = int64(transaction.Total)
			break
		}
	}
	return &JobTakeResponse{
		JobId:       job.ID.String(),
		Title:       job.Title,
		Description: job.Description,
		Reward:      rewardTotal,
	}
}
