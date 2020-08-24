package transaction

import (
	"github.com/maulIbra/clean-architecture-go/api-master/models"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
)

type TransactionUsecase struct {
	repo ITransactionRepo
}

func (t TransactionUsecase) PostTransaction(transaction *models.Transaction) error {
	transaction.TransactionDate = utils.GetTimeNow()
	user,err := t.repo.ReadUserByID(transaction.UserID)
	if err != nil {
		return err
	}
	balanceUser := user.Balance - transaction.Amount

	recipient,err := t.repo.ReadUserByID(transaction.RecipientID)
	if err != nil {
		return err
	}
	balanceRecepient := recipient.Balance + transaction.Amount

	log.Print(balanceRecepient)
	log.Print(balanceUser)
	log.Print(recipient)

	err = t.repo.PostTransaction(transaction,balanceUser,balanceRecepient)
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionUsecase) GetTransaction() ([]models.Transaction, error) {
	transactionList,err := t.repo.GetTransaction()
	if err != nil {
		return nil, err
	}
	return transactionList,nil
}

func (t TransactionUsecase) GetTransactionByID(userId int) ([]models.Transaction, error) {
	transactionList,err := t.repo.GetTransactionByID(userId)
	if err != nil {
		return nil, err
	}
	return transactionList,nil
}

func NewTransactionUsecase(repo ITransactionRepo) ITransactionUsecase{
	return &TransactionUsecase{
		repo: repo,
	}
}
