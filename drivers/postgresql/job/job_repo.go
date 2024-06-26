package job

import (
	"errors"
	"fmt"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/user"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/irvansn/go-find-helpers/utils"
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

	err := jobDb.Transactions[0].Payment.Create(job.ID, user.Email)
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

func (r *Repo) FindRelated(job *entities.Job, user *middlewares.Claims) error {
	jobDb := FromUseCase(job)

	if err := r.DB.Preload("Transactions.Payment").Preload(clause.Associations).Where("user_id = ?", user.ID).First(&jobDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*job = *jobDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(jobs *[]entities.Job, user *middlewares.Claims, statusFilter string, categoryIdFilter string) error {
	var jobsDb []Job

	query := r.DB.Preload(clause.Associations)
	if user.Role == "CUSTOMER" {
		query = query.Where("user_id = ?", user.ID)
	}

	if statusFilter != "" {
		query = query.Where("status = ?", statusFilter)
	}

	if categoryIdFilter != "" {
		query = query.Where("category_id = ?", categoryIdFilter)
	}

	if err := query.Find(&jobsDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	fmt.Println("adw", utils.PrettyPrint(jobsDb))

	for _, job := range jobsDb {
		*jobs = append(*jobs, *job.ToUseCase())
	}
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

func (r *Repo) Update(job *entities.Job, user *middlewares.Claims) error {
	jobDb := FromUseCase(job)

	if err := r.DB.Model(&jobDb).Where("id = ? AND user_id = ?", job.ID, user.ID).Updates(&jobDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
	}

	*job = *jobDb.ToUseCase()
	return nil
}

func (r *Repo) Delete(job *entities.Job, user *middlewares.Claims) error {
	jobDb := FromUseCase(job)

	db := r.DB.Where("id = ? AND user_id = ?", job.ID, user.ID).Delete(&jobDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
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

	fmt.Println(utils.PrettyPrint(jobDb.ID))
	fmt.Println(utils.PrettyPrint(jobDb.Status))

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

	// if unfilled helper slot refund to the CUSTOMER balance
	unfilledHelperSlot := int(jobDb.HelperRequired) - (len(jobDb.Transactions) - 1)
	customerUser := user.User{ID: jobDb.UserID}
	if err := tx.First(&customerUser).Error; err != nil {
		tx.Rollback()
	}
	customerUser.CurrentBalance += jobDb.RewardEarned * float64(unfilledHelperSlot)
	if err := tx.Save(&customerUser).Error; err != nil {
		tx.Rollback()
	}

	for i, _transaction := range job.Transactions {
		if _transaction.Type == "MONEY_IN" {
			job.Transactions[i].Payment.Status = "SUCCESS"

			helperUser := user.User{ID: _transaction.UserID}
			if err := tx.First(&helperUser).Error; err != nil {
				tx.Rollback()
			}

			if err := tx.Save(job.Transactions[i].Payment).Error; err != nil {
				tx.Rollback()
			}

			helperUser.CurrentBalance += _transaction.Total
			if err := tx.Save(&helperUser).Error; err != nil {
				tx.Rollback()
			}

			if err := tx.Model(&jobDb).Where("id = ?", jobDb.ID).Update("status", "DONE").Error; err != nil {
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

func (r *Repo) CustomerService(cs *entities.JobCustomerService, user *middlewares.Claims) error {
	csDb := CustomerServiceFromUseCase(cs)

	csResponse, err := csDb.Talk()
	if err != nil {
		return err
	}

	*cs = *csResponse
	return nil
}
