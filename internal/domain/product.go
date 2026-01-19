package domain

import "time"

type Product struct {
	ID        int     `gorm:"primaryKey"`   // ID do produto
	Name      string  `gorm:"not null"`     // Nome do produto (obrigatório)
	Category  string  `gorm:"not null"`     // Categoria do produto (obrigatório)
	Brand     string  `gorm:"not null"`     // Marca do produto (obrigatório)
	Price     float64 `gorm:"not null"`     // Preço do produto
	HasPromo  bool    `gorm:"default:true"` // Produto ativo (padrão true)
	CreatedAt time.Time
}
