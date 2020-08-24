package models

type Transaction struct {
	UserID   int `json:"userId"`
	TransactionDate string `json:"transactionDate"`
	Amount int `json:"amount"`
	BankName string `json:"bankName"`
	RecipientID int `json:"recipientId"`
	Description string `json:"description"`
	Status int `json:"status"`
}

type TransactionResponse struct {
	transactionList []Transaction
}
