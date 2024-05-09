package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type JobIdRequest struct {
	JobId string `json:"job_id"`
}

func (r *JobIdRequest) JobIdToEntities() *entities.Job {
	jobId, _ := uuid.Parse(r.JobId)

	return &entities.Job{
		ID: jobId,
	}
}
