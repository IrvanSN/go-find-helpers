package entities

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/middlewares"
	"time"
)

type Job struct {
	ID             uuid.UUID
	Title          string
	Description    string
	RewardEarned   float64
	FromAddress    Address
	ToAddress      Address
	Status         string
	HelperRequired uint
	Category       Category
	UserID         uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Transactions   []Transaction
	Thumbnails     []Thumbnail
}

type JobCustomerService struct {
	Question string
	Answer   string
}

type JobRepositoryInterface interface {
	Create(Job *Job, user *middlewares.Claims) error
	FindRelated(job *Job, user *middlewares.Claims) error
	Find(job *Job) error
	AddHelper(job *Job) error
	UpdateStatus(job *Job) error
	PaymentCallback(job *Job) error
	MarkAsDone(job *Job) error
	GetAll(job *[]Job, user *middlewares.Claims, statusFilter string, categoryIdFilter string) error
	Update(job *Job, user *middlewares.Claims) error
	Delete(job *Job, user *middlewares.Claims) error
	CustomerService(cs *JobCustomerService, user *middlewares.Claims) error
}

type JobUseCaseInterface interface {
	Create(job *Job, user *middlewares.Claims) (Job, error)
	Take(job *Job, user *middlewares.Claims) (Job, error)
	PaymentCallback(job *Job) (Job, error)
	MarkAsDone(job *Job, user *middlewares.Claims) (Job, error)
	MarkAsOnProgress(job *Job, user *middlewares.Claims) (Job, error)
	GetAll(job *[]Job, user *middlewares.Claims, statusFilter string, categoryIdFilter string) ([]Job, error)
	Update(job *Job, user *middlewares.Claims) (Job, error)
	Delete(job *Job, user *middlewares.Claims) (Job, error)
	GetJobDetail(job *Job, user *middlewares.Claims) (Job, error)
	CustomerService(cs *JobCustomerService, user *middlewares.Claims) (JobCustomerService, error)
}
