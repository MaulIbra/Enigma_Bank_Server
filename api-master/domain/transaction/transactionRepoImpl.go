package transaction

import (
	"database/sql"
	"github.com/maulIbra/clean-architecture-go/api-master/models"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
)

type TransactionRepo struct {
	db *sql.DB
}


func NewTransactionRepo(db *sql.DB) ITransactionRepo{
	return &TransactionRepo{
		db: db,
	}
}

func (t TransactionRepo) GetTransaction() ([]models.Transaction, error) {
	transactionList := []models.Transaction{}

	stmt, err := t.db.Prepare(utils.SELECT_TRANSACTION)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := models.Transaction{}
		err := rows.Scan(&p.UserID,&p.TransactionDate,&p.Amount,&p.BankName,&p.Description,&p.Status)
		if err != nil {
			return nil, err
		}
		transactionList = append(transactionList,p)
	}
	return transactionList,nil
}

func (t TransactionRepo) GetTransactionByID(userId int) ([]models.Transaction, error) {
	transactionList := []models.Transaction{}
	stmt, err := t.db.Prepare(utils.SELECT_TRANSACTION_BY_ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := models.Transaction{}
		err := rows.Scan(&p.UserID,&p.TransactionDate,&p.Amount,&p.BankName,&p.Description,&p.Status)
		if err != nil {
			return nil, err
		}
		transactionList = append(transactionList,p)
	}
	return transactionList,nil
}

func (t TransactionRepo) PostTransaction(transaction *models.Transaction,balanceUser,balanceRecipient int) error {
	tx, err := t.db.Begin()

	transfer := 1
	receive := 0

	transaction.Status = transfer

	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_TRANSACTION)
	defer stmt.Close()
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(transaction.UserID, transaction.TransactionDate, transaction.Amount,transaction.BankName,transaction.Description,transaction.Status); err != nil {
		tx.Rollback()
		return err
	}

	stmt2, err := tx.Prepare(utils.UPDATE_BALANCE_USER)
	defer stmt2.Close()
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}

	if _, err := stmt2.Exec(balanceUser,transaction.UserID); err != nil {
		tx.Rollback()
		return err
	}

	//transaction recepient
	stmt3, err := tx.Prepare(utils.INSERT_TRANSACTION)
	defer stmt3.Close()
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}

	if _, err := stmt3.Exec(transaction.RecipientID, transaction.TransactionDate, transaction.Amount,transaction.BankName,transaction.Description,receive); err != nil {
		tx.Rollback()
		return err
	}

	//update balance recepient
	stmt4, err := tx.Prepare(utils.UPDATE_BALANCE_USER)
	defer stmt4.Close()
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}

	if _, err := stmt4.Exec(balanceRecipient,transaction.RecipientID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (t TransactionRepo) ReadUserByID(id int) (*models.User, error) {
	stmt, err := t.db.Prepare(utils.SELECT_USER_BY_ID)
	u := models.User{}
	if err != nil {
		return &u, err
	}
	errQuery := stmt.QueryRow(id).Scan(&u.UserID, &u.Username, &u.Password,&u.Balance)

	if errQuery != nil {
		return &u, err
	}

	defer stmt.Close()
	return &u, nil
}