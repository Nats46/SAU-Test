package payload

import "time"

type ItemPayload struct {
	Id           int       `json:"id"`
	Name         string    `json:"Name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	CategoryId   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type CategoryPost struct {
	Name string `json:"name"`
}

type ItemPost struct {
	Id          int       `json:"id"`
	Name        string    `json:"Name"`
	CategoryId  int       `json:"category_id"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type ItemPut struct{
	Name        string    `json:"Name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}