package domain

import (
	"time"
)

type Product struct {
	ProductID   int    `gorm:"primaryKey" json:"product_id"`
	ProductName string `gorm:"varchar(40);not null" json:"product_name" binding:"required"`
	Type        string `gorm:"varchar(15);not null" json:"type" binding:"required"`
	Brand       string `gorm:"varchar(25);not null" json:"brand" binding:"required"`
	Color       string `gorm:"varchar(15);not null" json:"color" binding:"required"`
	Notes       string `gorm:"varchar(80);not null" json:"notes" binding:"required"`
}

type Post struct {
	ID       int       `gorm:"primaryKey" json:"id_post"`
	UserID   int       `json:"user_id" binding:"required"`
	Date     time.Time `json:"date" gorm:"type:date"`
	Product  Product   `json:"product" gorm:"embedded"`
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
