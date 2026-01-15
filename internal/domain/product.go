package model

// import "time"

// type Product struct {
// 	ID          string  `gorm:"primaryKey"`                  // ID do produto
// 	SellerID    string  `gorm:"foreignKey:ID;references:ID"` // ID do vendedor que está vendendo o produto (chave estrangeira)
// 	Name        string  `gorm:"not null"`                    // Nome do produto (obrigatório)
// 	Type        string  `gorm:"not null"`                    // Tipo do produto (obrigatório)
// 	Brand       string  `gorm:"not null"`                    // Marca do produto (obrigatório)
// 	Category    string  `gorm:"not null"`                    // Categoria do produto (obrigatório)
// 	Description string  `gorm:"not null"`                    // Descrição do produto (obrigatório)
// 	Price       float64 `gorm:"not null"`                    // Preço do produto
// 	Active      bool    `gorm:"default:true"`                // Produto ativo (padrão true)
// 	CreatedAt   time.Time
// }
