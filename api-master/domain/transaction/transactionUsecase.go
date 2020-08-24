package transaction

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type ITransactionUsecase interface {
	PostTransaction(transaction *models.Transaction) error
	GetTransaction() ([]models.Transaction, error)
	GetTransactionByID(userId int) ([]models.Transaction, error)
}
