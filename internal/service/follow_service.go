package service

import (
	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"github.com/JeffersonTupinamba/api-social-meli/internal/repository"
)

// FollowService é o serviço para as operações de follow (injeção de dependência)
type FollowService struct {
	FollowRepository *repository.FollowRepository
}

// US 0001: Poder "seguir" um vendedor específico
// FollowUserService é a função que permite a um usuário seguir um vendedor específico
func (s *FollowService) FollowUserService(u *domain.UserFollow) error {
	err := s.FollowRepository.FollowUser(u)
	if err != nil {
		return err
	}
	return nil
}
