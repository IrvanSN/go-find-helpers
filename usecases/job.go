package usecases

import "github.com/irvansn/go-find-helpers/entities"

type JobUseCase struct {
	repository entities.JobRepositoryInterface
}

func NewJobUseCase(repository entities.JobRepositoryInterface) *JobUseCase {
	return &JobUseCase{
		repository: repository,
	}
}

func (j *JobUseCase) Create(job *entities.Job) (entities.Job, error) {
	//if err := j.repository.Create(job); err != nil {
	//	return entities.Job{}, err
	//}
	return *job, nil
}
