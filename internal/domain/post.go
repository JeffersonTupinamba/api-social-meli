package domain

import (
	"time"
)

type Post struct {
	ID       int       `gorm:"primaryKey" json:"id_post"`
	UserID   int       `json:"user_id" binding:"required"`
	Date     time.Time `json:"date" gorm:"type:date"`
	Product  Product   `json:"product" gorm:"embedded;embeddedPrefix:product_"`
	Category int       `json:"category" binding:"required"`
	Price    float64   `json:"price" binding:"required"`
}

// Struct para receber o JSON do Postman (com data em string)
type RequestPostCreate struct {
	UserID   int     `json:"user_id" binding:"required"`
	Date     string  `json:"date" binding:"required"` // Formato "DD-MM-YYYY"
	Product  Product `json:"product" binding:"required"`
	Category int     `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
}

type ResponsePostCreate struct {
	ID       int       `json:"id_post"`
	UserID   int       `json:"user_id"`
	Date     time.Time `json:"date"`
	Product  Product   `json:"product"`
	Category int       `json:"category"`
	Price    float64   `json:"price"`
}
