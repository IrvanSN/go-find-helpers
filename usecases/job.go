package usecases

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
)

type JobUseCase struct {
	repository entities.JobRepositoryInterface
}

func NewJobUseCase(repository entities.JobRepositoryInterface) *JobUseCase {
	return &JobUseCase{
		repository: repository,
	}
}

func (j *JobUseCase) Create(job *entities.Job, user *middlewares.Claims) (entities.Job, error) {
	if user.Role != "CUSTOMER" {
		return entities.Job{}, constant.ErrNotAuthorized
	}

	if job.Title == "" || job.Description == "" {
		return entities.Job{}, constant.ErrEmptyInput
	}

	if job.RewardEarned < 10000 || job.HelperRequired < 1 {
		return entities.Job{}, constant.ErrInvalidRequest
	}

	job.ID = uuid.New()
	job.Status = "CLOSED"
	job.UserID = user.ID

	transaction := entities.Transaction{}
	transaction.ID = uuid.New()
	transaction.Payment.ID = uuid.New()
	transaction.Payment.TransactionID = transaction.ID

	subTotal := job.RewardEarned * float64(job.HelperRequired)
	tax := (subTotal / 100) * 5

	transaction.Type = "MONEY_OUT"
	transaction.Payment.Status = "PENDING"
	transaction.UserID = user.ID
	transaction.JobID = job.ID
	transaction.SubTotal = subTotal
	transaction.Tax = tax
	transaction.Total = subTotal + tax

	transaction.Payment.Amount = int64(transaction.Total)

	job.Transactions = append(job.Transactions, transaction)

	if err := j.repository.Create(job, user); err != nil {
		return entities.Job{}, err
	}

	return *job, nil
}

func (j *JobUseCase) Take(job *entities.Job, user *middlewares.Claims) (entities.Job, error) {
	if job.ID == uuid.Nil {
		return entities.Job{}, constant.ErrEmptyInput
	}

	job.UserID = user.ID

	if user.Role != "HELPER" {
		return entities.Job{}, constant.ErrNotAuthorized
	}

	if err := j.repository.Find(job); err != nil {
		return entities.Job{}, err
	}

	if job.Status == "ON_PROGRESS" || job.Status == "CLOSED" {
		return entities.Job{}, constant.ErrJobAlreadyClosed
	}

	if job.Status == "DONE" {
		return entities.Job{}, constant.ErrJobAlreadyDone
	}

	helperAlreadyTakeTheJob := false
	for _, transaction := range job.Transactions {
		if transaction.UserID == user.ID && transaction.Type == "MONEY_IN" {
			helperAlreadyTakeTheJob = true
			break
		}
	}

	if helperAlreadyTakeTheJob {
		return entities.Job{}, constant.ErrHelperAlreadyTakeTheJob
	}

	if (len(job.Transactions) + 1) > (int(job.HelperRequired) + 1) {
		return entities.Job{}, constant.ErrJobAlreadyFull
	}

	if (int(job.HelperRequired) + 1) == (len(job.Transactions) + 1) {
		job.Status = "ON_PROGRESS"
		if err := j.repository.UpdateStatus(job); err != nil {
			return entities.Job{}, constant.ErrFailedUpdate
		}
	}

	var transaction entities.Transaction
	transaction.ID = uuid.New()
	transaction.Type = "MONEY_IN"
	transaction.Payment.Status = "PENDING"
	transaction.UserID = user.ID
	transaction.JobID = job.ID

	transaction.SubTotal = job.RewardEarned
	tax := (transaction.SubTotal / 100) * 5
	transaction.Tax = tax
	transaction.Total = transaction.SubTotal - tax

	transaction.Payment.ID = uuid.New()
	transaction.Payment.TransactionID = transaction.ID
	transaction.Payment.Amount = int64(transaction.Total)

	var newJob entities.Job
	newJob.ID = job.ID
	newJob.UserID = user.ID
	newJob.Transactions = append(newJob.Transactions, transaction)

	if err := j.repository.AddHelper(&newJob); err != nil {
		return entities.Job{}, err
	}

	job.Transactions = append(job.Transactions, transaction)

	return *job, nil
}

func (j *JobUseCase) PaymentCallback(job *entities.Job) (entities.Job, error) {
	paymentId := job.Transactions[0].Payment.ID
	paymentStatus := job.Transactions[0].Payment.Status

	if err := j.repository.Find(job); err != nil {
		return entities.Job{}, err
	}

	job.Transactions[0].Payment.ID = paymentId

	if paymentStatus == "PAID" {
		job.Status = "OPEN"
		job.Transactions[0].Payment.Status = "SUCCESS"
	}

	if paymentStatus == "EXPIRED" {
		job.Status = "CLOSED"
		job.Transactions[0].Payment.Status = "EXPIRED"
	}

	if err := j.repository.PaymentCallback(job); err != nil {
		return entities.Job{}, err
	}

	return *job, nil
}

func (j *JobUseCase) MarkAsDone(job *entities.Job, user *middlewares.Claims) (entities.Job, error) {
	if user.Role != "CUSTOMER" {
		return entities.Job{}, constant.ErrNotAuthorized
	}

	if job.ID == uuid.Nil {
		return entities.Job{}, constant.ErrEmptyInput
	}

	if err := j.repository.FindRelated(job, user); err != nil {
		return entities.Job{}, err
	}

	if job.Status == "CLOSED" {
		return entities.Job{}, constant.ErrJobAlreadyClosed
	}

	if job.Status == "DONE" {
		return entities.Job{}, constant.ErrJobAlreadyDone
	}

	if job.Status == "OPEN" {
		return entities.Job{}, constant.ErrJobStillOpened
	}

	if err := j.repository.MarkAsDone(job); err != nil {
		return entities.Job{}, err
	}

	if err := j.repository.FindRelated(job, user); err != nil {
		return entities.Job{}, err
	}

	return *job, nil
}

func (j *JobUseCase) MarkAsOnProgress(job *entities.Job, user *middlewares.Claims) (entities.Job, error) {
	if user.Role != "CUSTOMER" {
		return entities.Job{}, constant.ErrNotAuthorized
	}

	if job.ID == uuid.Nil {
		return entities.Job{}, constant.ErrEmptyInput
	}

	if err := j.repository.FindRelated(job, user); err != nil {
		return entities.Job{}, err
	}

	if job.Status == "CLOSED" {
		return entities.Job{}, constant.ErrJobAlreadyClosed
	}

	if job.Status == "DONE" {
		return entities.Job{}, constant.ErrJobAlreadyDone
	}

	if job.Status == "ON_PROGRESS" {
		return entities.Job{}, constant.ErrJobAlreadyOnProgress
	}

	job.Status = "ON_PROGRESS"

	if err := j.repository.UpdateStatus(job); err != nil {
		return entities.Job{}, constant.ErrFailedUpdate
	}

	if err := j.repository.FindRelated(job, user); err != nil {
		return entities.Job{}, err
	}

	return *job, nil
}

func (j *JobUseCase) GetAll(job *[]entities.Job, user *middlewares.Claims, status string) ([]entities.Job, error) {
	if err := j.repository.GetAll(job, user, status); err != nil {
		return []entities.Job{}, err
	}

	return *job, nil
}

func (j *JobUseCase) Update(job *entities.Job, user *middlewares.Claims) (entities.Job, error) {
	if user.Role != "CUSTOMER" {
		return entities.Job{}, constant.ErrNotAuthorized
	}

	if err := j.repository.Update(job, user); err != nil {
		return entities.Job{}, err
	}

	if err := j.repository.FindRelated(job, user); err != nil {
		return entities.Job{}, err
	}

	return *job, nil
}

func (j *JobUseCase) Delete(job *entities.Job, user *middlewares.Claims) (entities.Job, error) {
	if user.Role != "CUSTOMER" {
		return entities.Job{}, constant.ErrNotAuthorized
	}

	if err := j.repository.Delete(job, user); err != nil {
		return entities.Job{}, err
	}

	return *job, nil
}
