package job

import (
	"errors"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo struct {
	DB *gorm.DB
}

func NewJobRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(job *entities.Job, user *middlewares.Claims) error {
	jobDb := FromUseCase(job)

	err := jobDb.Transactions[0].Payment.Create(user.Email)
	if err != nil {
		return err
	}

	if err := r.DB.Omit("FromAddress").Omit("ToAddress").Create(&jobDb).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return constant.ErrDuplicatedData
		}
		return err
	}

	*job = *jobDb.ToUseCase()
	return nil
}

func (r *Repo) Find(job *entities.Job) error {
	jobDb := FromUseCase(job)

	if err := r.DB.Preload(clause.Associations).First(&jobDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*job = *jobDb.ToUseCase()
	return nil
}

func (r *Repo) AddHelper(job *entities.Job) error {
	jobDb := FromUseCase(job)

	if err := r.DB.Model(&jobDb).Association("Transactions").Append(&jobDb.Transactions); err != nil {
		return constant.ErrInsertDatabase
	}

	if err := r.DB.Create(&jobDb.Transactions[0].Payment).Error; err != nil {
		return constant.ErrInsertDatabase
	}

	*job = *jobDb.ToUseCase()
	return nil
}
