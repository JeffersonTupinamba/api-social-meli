package domain

import "time"

type UserFollow struct {
	ID         int       `gorm:"primaryKey"`
	FollowerID int       `gorm:"not null;uniqueIndex:idx_follow"` // quem segue (customer)
	SellerID   int       `gorm:"not null;uniqueIndex:idx_follow"` // quem é seguido (seller)
	Date       time.Time `gorm:"autoCreateTime"`                  // Data de criação do follow
}
