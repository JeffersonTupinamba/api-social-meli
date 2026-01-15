package model

import "time"

// type Post struct {
// 	ID        string    `gorm:"primaryKey"`
// 	SellerID  string    `gorm:"foreignKey:ID;references:ID"` // ID do vendedor que está vendendo o produto (chave estrangeira)
// 	ProductID string    `gorm:"foreignKey:ID;references:ID"` // ID do produto que está sendo vendido (chave estrangeira)
// 	Content   string    `gorm:"not null"`                    // Conteúdo do post (obrigatório)
// 	CreatedAt time.Time `gorm:"autoCreateTime"`              // Data de criação do post
// }

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;index"` // vendedor que criou o post
	Content   string `gorm:"not null"`       // ou Title/Description etc
	CreatedAt time.Time
}
