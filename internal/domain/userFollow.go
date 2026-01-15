package model

// import "time"

// type UserFollow struct {
// 	ID         uint `gorm:"primaryKey"`
// 	FollowerID uint `gorm:"not null;index"` // quem segue (customer)
// 	FollowedID uint `gorm:"not null;index"` // quem é seguido (seller)
// 	CreatedAt  time.Time

// 	// chave estrangeira para o usuário que está seguindo e o usuário que está sendo seguido
// 	Follower User `gorm:"foreignKey:FollowerID"`
// 	Followed User `gorm:"foreignKey:FollowedID"`
// }
