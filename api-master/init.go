package api_master

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/api-master/domain/authentication"
	"github.com/maulIbra/clean-architecture-go/api-master/domain/transaction"
)

func Init(router *mux.Router, db *sql.DB) {

	//transaction
	transactionRepo := transaction.NewTransactionRepo(db)
	transactionUsecase := transaction.NewTransactionUsecase(transactionRepo)
	transactionController := transaction.NewTransactionController(transactionUsecase)
	transactionController.Transaction(router)

	//authentication
	AuthenticationRepo := authentication.NewAuthenticationRepo(db)
	AuthenticationUsecase := authentication.NewAuthenticationUsecase(AuthenticationRepo)
	AuthenticationController := authentication.NewAuthenticationController(AuthenticationUsecase)
	AuthenticationController.Authenticate(router)
}
