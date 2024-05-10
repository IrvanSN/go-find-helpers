package usecases

import (
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
)

type TransactionUseCase struct {
	repository entities.TransactionRepositoryInterface
}

func NewTransactionUseCase(repository entities.TransactionRepositoryInterface) *TransactionUseCase {
	return &TransactionUseCase{repository: repository}
}

func (t *TransactionUseCase) GetAllTransactions(transactions *[]entities.Transaction, transactionType string, user *middlewares.Claims) ([]entities.Transaction, error) {
	if user.Role != "ADMIN" {
		return []entities.Transaction{}, constant.ErrNotAuthorized
	}

	if err := t.repository.GetAllTransaction(transactions, transactionType); err != nil {
		return []entities.Transaction{}, err
	}

	return *transactions, nil
}
