package model

type Services struct {
	Id          string `json:"id"`
	ServiceName string `json:"service_name"`
	Unit        string `json:"unit"`
	Price       int    `json:"price"`
}
