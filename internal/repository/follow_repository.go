package repository

import (
	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"gorm.io/gorm"
)

// FollowRepository é o repositório para as operações de follow (injeção de dependência)
type FollowRepository struct {
	DB *gorm.DB
}

func (r *FollowRepository) CheckFollowByUserIdAndSellerId(userId, sellerId int) (bool, error) {
	var count int64

	resp := r.DB.
		Model(&domain.UserFollow{}).
		Where("follower_id = ? AND seller_id = ?", userId, sellerId).
		Count(&count)

	if resp.Error != nil {
		return false, resp.Error
	}

	return count > 0, nil
}

// Cria um novo follow
func (r *FollowRepository) CreateFollow(f *domain.UserFollow) error {
	resp := r.DB.Create(f)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

// busca um vendedor pelo ID
func (r *FollowRepository) GetSellerById(id int) (*domain.User, error) {
	var seller domain.User
	resp := r.DB.Model(&domain.User{}).First(&seller, id)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &seller, nil
}
