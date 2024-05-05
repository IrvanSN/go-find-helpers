package response

import "github.com/irvansn/go-find-helpers/entities"

type JobTakeResponse struct {
	JobId       string `json:"job_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Reward      int64  `json:"reward"`
}

func TakeResponseFromUseCase(job *entities.Job) *JobTakeResponse {
	return &JobTakeResponse{
		JobId:       job.ID.String(),
		Title:       job.Title,
		Description: job.Description,
		// Todo: handle payment.amount transactions created by helper user
		Reward: int64(job.RewardEarned),
	}
}
