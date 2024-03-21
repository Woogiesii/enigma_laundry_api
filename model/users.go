package model

import "time"

type Users struct {
	Id           string    `json:"id"`
	CustomerName string    `json:"customer_name"`
	PhoneNumber  string    `json:"phone_number"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	DateCreated  time.Time `json:"date_created"`
}

type UsersData struct {
	Id           string    `json:"id"`
	CustomerName string    `json:"customer_name"`
	PhoneNumber  string    `json:"phone_number"`
	Username     string    `json:"username"`
	Role         string    `json:"role"`
	DateCreated  time.Time `json:"date_created"`
}
