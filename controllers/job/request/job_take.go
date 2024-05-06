package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type JobTakeRequest struct {
	JobId string `json:"job_id"`
}

func (j *JobTakeRequest) JobTakeToEntities() *entities.Job {
	jobId, _ := uuid.Parse(j.JobId)

	return &entities.Job{
		ID: jobId,
	}
}
