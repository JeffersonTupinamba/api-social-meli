package repository

import (
	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"gorm.io/gorm"
)

// FollowRepository é o repositório para as operações de follow (injeção de dependência)
type FollowRepository struct {
	DB *gorm.DB
}

// US 0001: Poder "seguir" um vendedor específico
// FollowUser é a função que permite a um usuário seguir um vendedor específico
func (r *FollowRepository) FollowUser(u *domain.UserFollow) error {
	resp := r.DB.Create(&u)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
