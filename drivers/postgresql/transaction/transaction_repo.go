package transaction

import (
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) GetAllTransaction(transactions *[]entities.Transaction, transactionType string) error {
	var transactionDb []Transaction

	db := r.DB.Preload(clause.Associations)

	if transactionType != "" {
		db = db.Where("type = ?", transactionType)
	}

	if err := db.Find(&transactionDb).Error; err != nil {
		return err
	}

	for _, _transaction := range transactionDb {
		*transactions = append(*transactions, *_transaction.ToUseCase())
	}
	return nil
}
