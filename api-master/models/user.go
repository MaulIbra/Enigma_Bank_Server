package models

type User struct {
	UserID   int `json:"userId"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Balance int `json:"balance"`
	Token Token		`json:"authentication"`
}
