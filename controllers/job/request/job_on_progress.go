package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type JobOnProgressRequest struct {
	JobId string `json:"job_id"`
}

func (r *JobOnProgressRequest) JobOnProgressToEntities() *entities.Job {
	jobId, _ := uuid.Parse(r.JobId)

	return &entities.Job{
		ID: jobId,
	}
}
