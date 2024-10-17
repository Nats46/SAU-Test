package model

import "time"

type Item struct {
	Id          int
	CategoryId  int
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
}