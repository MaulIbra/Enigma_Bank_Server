package transaction

import (
	"github.com/maulIbra/clean-architecture-go/api-master/models"
)

type ITransactionRepo interface {
	PostTransaction(transaction *models.Transaction,balanceUser,balanceRecepient int) error
	GetTransaction() ([]models.Transaction, error)
	GetTransactionByID(userId int) ([]models.Transaction, error)
	ReadUserByID(id int) (*models.User, error)
}
