package repository

import (
	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"gorm.io/gorm"
)

// UserRepository é o repositório para as operações de usuário (injeção de dependência)
type UserRepository struct {
	DB *gorm.DB // banco de dados para realizar as operações no banco de dados
}

// CRIA UM NOVO USUÁRIO
// CreateUser é a função que cria um novo usuário no banco de dados
func (r *UserRepository) CreateUser(u *domain.User) error {
	resp := r.DB.Create(&u)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

// Lista todos os usuários
// ListUsers é a função que lista todos os usuários no banco de dados
func (r *UserRepository) ListUsers(l *[]domain.UserListResponse) error {
	resp := r.DB.Model(&domain.User{}).Select("id, name, email, role").Scan(l)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
