package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Weight          float32 `json:"weight"`
	Length          float32 `json:"length"`
	Width           float32 `json:"width"`
	Height          float32 `json:"height"`
	Price           float32 `json:"price"`
	AdditionalPrice float32 `json:"additionalPrice"`
	QTY             int     `json:"QTY"`
	UserID          uint
	StoreID         uint
	CategoriesID    uint
	CurrencyID      uint
	OrderItems      []OrderItems
	Reviews         []Reviews
	CartItems       []CartItems
}
