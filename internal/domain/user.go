package domain

import "time"

// User é a representação da nossa tabela no banco de dados
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;uniqueIndex"`
	Role      string    `json:"role" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// RequestUserCreate define o que esperamos receber no POST /users
type RequestUserCreate struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Role      string    `json:"role" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// RequestUserUpdate define o que permitimos atualizar no PUT /users/:id
type RequestUserUpdate struct {
	Name      string    `json:"name"`
	Email     string    `json:"email" binding:"omitempty,email"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// RequestUserFollow define o que esperamos receber no POST /users/:userId/follow/:sellerId
type RequestUserFollow struct {
	SellerID int `json:"seller_id" binding:"required"`
}

type ResquestUserUnfollow struct {
	SellerID int `json:"seller_id" binding:"required"`
}
