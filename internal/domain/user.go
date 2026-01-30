package domain

import "time"

// User é a representação da nossa tabela no banco de dados
type User struct {
	ID    int       `gorm:"primaryKey"`
	Name  string    `gorm:"varchar(15);not null"`
	Email string    `gorm:"varchar(40);not null;uniqueIndex"`
	Role  string    `gorm:"varchar(15);not null"`
	Date  time.Time `gorm:"autoCreateTime"`
}

// RequestUserCreate define o que esperamos receber no POST /users
type RequestUserCreate struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required"`
}

// ResponseUser define o que esperamos retornar no GET /users/:id
type ResponseUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Date      time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// RequestUserUpdate define o que permitimos atualizar no PUT /users/:id
type RequestUserUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
	Role  string `json:"role"`
}

type FollowerDTO struct {
	UserID   int    `json:"userId"`
	UserName string `json:"userName"`
}

type UserFollowersListResponse struct {
	UserID    int           `json:"userId"`
	UserName  string        `json:"userName"`
	Followers []FollowerDTO `json:"followers"`
}

type UserFollowedListResponse struct {
	UserID          int           `json:"userId"`
	UserName        string        `json:"userName"`
	FollowedSellers []FollowerDTO `json:"followedSellers"`
}

type UserFollowersCountResponse struct {
	UserID         int    `json:"user_id"`
	UserName       string `json:"user_name"`
	FollowersCount int64  `json:"followers_count"`
}

type UserListResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
