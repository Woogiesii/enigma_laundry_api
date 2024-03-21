package dto

import "time"

type UsersRequestDto struct {
	Id           string    `json:"id"`
	CustomerName string    `json:"customer_name"`
	PhoneNumber  string    `json:"phone_number"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	DateCreated  time.Time `json:"date_created"`
}

type LoginRequestDto struct {
	Username string `json:"username" binding:"required"`
	Pass     string `json:"password" binding:"required"`
}

type LoginResponseDto struct {
	AccesToken string `json:"accesToken"`
	UserId     string `json:"userId"`
}
