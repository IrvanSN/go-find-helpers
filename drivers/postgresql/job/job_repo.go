package job

import (
	"errors"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/user"
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

	if err := r.DB.Preload("Transactions.Payment").Preload(clause.Associations).First(&jobDb).Error; err != nil {
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

func (r *Repo) Update(job *entities.Job) error {
	jobDb := FromUseCase(job)

	if err := r.DB.Save(&jobDb).Error; err != nil {
		return constant.ErrInsertDatabase
	}

	*job = *jobDb.ToUseCase()
	return nil
}

func (r *Repo) UpdateStatus(job *entities.Job) error {
	jobDb := FromUseCase(job)

	if err := r.DB.Model(&jobDb).Where("id = ?", jobDb.ID).Update("status", jobDb.Status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*job = *jobDb.ToUseCase()
	return nil
}

func (r *Repo) PaymentCallback(job *entities.Job) error {
	jobDb := FromUseCase(job)

	if err := r.DB.Model(&jobDb.Transactions[0].Payment).Where("id = ?", jobDb.Transactions[0].Payment.ID).Update("status", jobDb.Transactions[0].Payment.Status).Error; err != nil {
		return constant.ErrFailedUpdate
	}

	if err := r.DB.Model(&jobDb).Where("id = ?", jobDb.ID).Update("status", jobDb.Status).Error; err != nil {
		return constant.ErrFailedUpdate
	}

	*job = *jobDb.ToUseCase()
	return nil
}

func (r *Repo) MarkAsDone(job *entities.Job) error {
	jobDb := FromUseCase(job)

	tx := r.DB.Begin()
	defer func() {
		if re := recover(); re != nil {
			tx.Rollback()
		}
	}()

	for i, _transaction := range job.Transactions {
		if _transaction.Type == "MONEY_IN" {
			job.Transactions[i].Payment.Status = "SUCCESS"

			newUser := user.User{ID: _transaction.UserID}
			if err := tx.First(&newUser).Error; err != nil {
				tx.Rollback()
			}

			if err := tx.Save(job.Transactions[i].Payment).Error; err != nil {
				tx.Rollback()
			}

			newUser.CurrentBalance += _transaction.Total
			if err := tx.Save(&newUser).Error; err != nil {
				tx.Rollback()
			}

			if err := tx.Model(&jobDb).Where("id = ?", jobDb.ID).Update("status", "CLOSED").Error; err != nil {
				tx.Rollback()
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	*job = *jobDb.ToUseCase()
	return nil
}
