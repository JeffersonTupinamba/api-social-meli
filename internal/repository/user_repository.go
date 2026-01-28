package repository

import (
	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB // banco de dados para realizar as operações no banco de dados (injeção de dependência) (é uma instância do banco de dados)
}

func (r *UserRepository) CreateUser(u *domain.User) error {
	resp := r.DB.Create(&u)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
