package model

type Transaction struct {
	Id             string `json:"id"`
	Users          string `json:"id_users"`
	Services       string `json:"id_services"`
	TransactionIn  int    `json:"transaction_in"`
	TransactionOut int    `json:"transaction_out"`
	Amount         int    `json:"amount"`
	CreatedAt      int    `json:"created_at"`
	UpdatedAt      int    `json:"updated_at"`
}

type Transactionnotepochdate struct {
	Id             string `json:"id"`
	Users          string `json:"id_users"`
	Services       string `json:"id_services"`
	TransactionIn  string `json:"transaction_in"`
	TransactionOut string `json:"transaction_out"`
	Amount         int    `json:"amount"`
	CreatedAt      int    `json:"created_at"`
	UpdatedAt      int    `json:"updated_at"`
}
