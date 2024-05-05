package job

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/address"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/category"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/payment"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/thumbnail"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/transaction"
	"github.com/irvansn/go-find-helpers/entities"
	"time"
)

type Job struct {
	ID             uuid.UUID `gorm:"type:varchar(100);"`
	Title          string    `gorm:"type:varchar(255);not null"`
	Description    string    `gorm:"type:text;not null"`
	RewardEarned   float64   `gorm:"type:decimal;not null"`
	FromAddressID  uuid.UUID `gorm:"type:varchar(100)"`
	FromAddress    address.Address
	ToAddressID    uuid.UUID `gorm:"type:varchar(100)"`
	ToAddress      address.Address
	Status         string    `gorm:"type:varchar(50);not null"`
	HelperRequired uint      `gorm:"not null"`
	CategoryID     uuid.UUID `gorm:"type:varchar(100);not null"`
	Category       category.Category
	UserID         uuid.UUID `gorm:"type:varchar(100);not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	Transactions   []transaction.Transaction
	Thumbnails     []thumbnail.Thumbnail
}

func FromUseCase(job *entities.Job) *Job {
	jobTransactions := make([]transaction.Transaction, len(job.Transactions))
	for i, _transaction := range job.Transactions {
		jobTransactions[i] = transaction.Transaction{
			ID:     _transaction.ID,
			UserID: _transaction.User.ID,
			Type:   _transaction.Type,
			Payment: payment.Payment{
				ID:            _transaction.Payment.ID,
				ExternalID:    _transaction.Payment.ExternalID,
				TransactionID: _transaction.Payment.TransactionID,
				Amount:        _transaction.Payment.Amount,
				InvoiceURL:    _transaction.Payment.InvoiceURL,
			},
			CreatedAt: _transaction.CreatedAt,
			UpdatedAt: _transaction.UpdatedAt,
		}
	}
	jobThumbnails := make([]thumbnail.Thumbnail, len(job.Thumbnails))
	for i, _thumbnail := range job.Thumbnails {
		jobThumbnails[i] = thumbnail.Thumbnail{
			ID:          uuid.New(),
			ImageKey:    _thumbnail.ImageKey,
			Description: _thumbnail.Description,
		}
	}

	return &Job{
		ID:             job.ID,
		Title:          job.Title,
		Description:    job.Description,
		RewardEarned:   job.RewardEarned,
		UserID:         job.User.ID,
		FromAddressID:  job.FromAddress.ID,
		FromAddress:    address.Address{ID: job.FromAddress.ID},
		ToAddressID:    job.ToAddress.ID,
		ToAddress:      address.Address{ID: job.ToAddress.ID},
		Status:         job.Status,
		HelperRequired: job.HelperRequired,
		Category:       category.Category{ID: job.Category.ID},
		CreatedAt:      job.CreatedAt,
		UpdatedAt:      job.UpdatedAt,
		Transactions:   jobTransactions,
		Thumbnails:     jobThumbnails,
	}
}

func (j *Job) ToUseCase() *entities.Job {
	jobTransactions := make([]entities.Transaction, len(j.Transactions))
	for i, _transaction := range j.Transactions {
		jobTransactions[i] = entities.Transaction{
			ID:   _transaction.ID,
			Type: _transaction.Type,
			Payment: entities.Payment{
				ID:            _transaction.Payment.ID,
				ExternalID:    _transaction.Payment.ExternalID,
				TransactionID: _transaction.Payment.TransactionID,
				Amount:        _transaction.Payment.Amount,
				InvoiceURL:    _transaction.Payment.InvoiceURL,
			},
			CreatedAt: _transaction.CreatedAt,
			UpdatedAt: _transaction.UpdatedAt,
		}
	}
	jobThumbnails := make([]entities.Thumbnail, len(j.Thumbnails))
	for i, _thumbnail := range j.Thumbnails {
		jobThumbnails[i] = entities.Thumbnail{
			ID:          _thumbnail.ID,
			ImageKey:    _thumbnail.ImageKey,
			Description: _thumbnail.Description,
		}
	}

	return &entities.Job{
		ID:             j.ID,
		Title:          j.Title,
		Description:    j.Description,
		RewardEarned:   j.RewardEarned,
		FromAddress:    entities.Address{ID: j.FromAddress.ID},
		ToAddress:      entities.Address{ID: j.ToAddress.ID},
		Status:         j.Status,
		HelperRequired: j.HelperRequired,
		Category:       entities.Category{ID: j.Category.ID},
		CreatedAt:      j.CreatedAt,
		UpdatedAt:      j.UpdatedAt,
		Transactions:   jobTransactions,
		Thumbnails:     jobThumbnails,
	}
}
