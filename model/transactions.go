package model

import "time"

type Transaction struct {
	Id             string    `json:"id"`
	Users          string    `json:"id_users"`
	Services       string    `json:"id_services"`
	TransactionIn  time.Time `json:"transaction_in"`
	TransactionOut time.Time `json:"transaction_out"`
	Amount         int       `json:"amount"`
	CreatedAt      int       `json:"created_at"`
	UpdatedAt      int       `json:"updated_at"`
}
