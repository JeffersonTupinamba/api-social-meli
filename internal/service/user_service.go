package service

import (
	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"github.com/JeffersonTupinamba/api-social-meli/internal/repository"
)

// UserHandler é o handler para as operações de usuário (injeção de dependência)
type UserService struct {
	UserRepository *repository.UserRepository
}

// CRIA UM NOVO USUÁRIO
// CreateUserService é a função que cria um novo usuário no banco de dados
func (s *UserService) CreateUserService(u *domain.User) error {
	err := s.UserRepository.CreateUser(u)
	if err != nil {
		return err
	}
	return nil
}
