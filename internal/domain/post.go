package domain

import "time"

type PostProduct struct {
	ID        int     `gorm:"primaryKey"`   // ID do produto
	Name      string  `gorm:"not null"`     // Nome do produto (obrigatório)
	Category  string  `gorm:"not null"`     // Categoria do produto (obrigatório)
	Brand     string  `gorm:"not null"`     // Marca do produto (obrigatório)
	Price     float64 `gorm:"not null"`     // Preço do produto
	HasPromo  bool    `gorm:"default:true"` // Produto ativo (padrão true)
	CreatedAt time.Time
}

type Post struct {
	ID     int       `gorm:"primaryKey"`
	UserID int       `gorm:"foreignKey:ID;references:ID"` // ID do vendedor que está vendendo o produto (chave estrangeira)
	Date   time.Time `gorm:"date"`                        // Data de criação do post
	Detail Product   `gorm:"detail"`                      // Detalhes do produto
}
