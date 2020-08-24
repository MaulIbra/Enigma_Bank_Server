package transaction

import (
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/api-master/models"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
	"net/http"
	"strconv"
)

type transactionController struct {
	usecase ITransactionUsecase
}

func NewTransactionController(usecase ITransactionUsecase) *transactionController{
	return &transactionController{usecase: usecase}
}

func (th *transactionController) Transaction(r *mux.Router) {
	transaction := r.PathPrefix("/transaction").Subrouter()
	//transaction.Use(middleware.TokenValidationMiddleware)
	transaction.HandleFunc("", th.readTransaction).Methods(http.MethodGet)
	transaction.HandleFunc("/{id}", th.readTransactionById).Methods(http.MethodGet)
	transaction.HandleFunc("", th.addTransaction).Methods(http.MethodPost)
}

func (th *transactionController) readTransaction(w http.ResponseWriter, r *http.Request){
	transactionList, err := th.usecase.GetTransaction()
	if err != nil {
		log.Print(err)
		utils.HandleResponseError(w, http.StatusBadGateway,utils.BAD_GATEWAY)
	}else {
		utils.HandleResponse(w, http.StatusOK, transactionList)
	}
}

func (th *transactionController) readTransactionById(writer http.ResponseWriter, request *http.Request) {
	id,_ := strconv.Atoi(utils.DecodePathVariabel("id", request))
	transactionList, err := th.usecase.GetTransactionByID(id)
	if err != nil {
		log.Print(err)
		utils.HandleResponseError(writer, http.StatusBadGateway,utils.BAD_GATEWAY)
	}else {
		utils.HandleResponse(writer, http.StatusOK, transactionList)
	}
}


func (th *transactionController) addTransaction(w http.ResponseWriter, r *http.Request){
	var transaction models.Transaction
	err := utils.JsonDecoder(&transaction,r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
	err = th.usecase.PostTransaction(&transaction)
	if err != nil {
		log.Print(err)
		utils.HandleResponseError(w, http.StatusBadGateway,utils.BAD_GATEWAY)
	}else{
		utils.HandleResponse(w, http.StatusOK,transaction)
	}
}

