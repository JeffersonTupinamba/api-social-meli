package domain

type Product struct {
	ProductID   int    `gorm:"primaryKey" json:"product_id"`
	ProductName string `gorm:"varchar(40);not null" json:"product_name" binding:"required"`
	Type        string `gorm:"varchar(15);not null" json:"type" binding:"required"`
	Brand       string `gorm:"varchar(25);not null" json:"brand" binding:"required"`
	Color       string `gorm:"varchar(15);not null" json:"color" binding:"required"`
	Notes       string `gorm:"varchar(80);not null" json:"notes" binding:"required"`
}
