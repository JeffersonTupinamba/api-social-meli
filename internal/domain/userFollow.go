package domain

import "time"

type UserFollow struct {
	ID         int       `gorm:"primaryKey"`
	FollowerID int       `gorm:"not null;index"` // quem segue (customer)
	SellerID   int       `gorm:"not null;index"` // quem é seguido (seller)
	Date       time.Time `gorm:"date"`           // Data de criação do follow
}
