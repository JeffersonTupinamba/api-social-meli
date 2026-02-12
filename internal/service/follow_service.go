package service

import (
	"errors"

	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"github.com/JeffersonTupinamba/api-social-meli/internal/repository"
)

// FollowService é o serviço para as operações de follow (injeção de dependência)
type FollowService struct {
	FollowRepository *repository.FollowRepository
	UserRepository   *repository.UserRepository
}

// US 0001: Poder "seguir" um vendedor específico
// FollowUserService é a função que permite a um usuário seguir um vendedor específico
func (s *FollowService) FollowUserService(userId, sellerId int) (*domain.UserFollow, error) {

	// verifica se o usuário está tentando seguir ele mesmo
	if userId == sellerId {
		return nil, errors.New("Você não pode seguir você mesmo.")
	}
	// verifica se o vendedor existe
	seller, err := s.UserRepository.GetUserById(sellerId)
	if err != nil {
		return nil, errors.New("Vendedor não encontrado.")
	}
	if seller.Role != "seller" {
		return nil, errors.New("O ID informado não pertence a um vendedor.")
	}
	// verifica se o usuário existe
	user, err := s.UserRepository.GetUserById(userId)
	if err != nil {
		return nil, errors.New("Usuário não encontrado.")
	}
	if user.Role != "customer" {
		return nil, errors.New("O ID informado não pertence a um comprador.")
	}
	// verifica se o usuário já está seguindo o vendedor
	exist, err := s.FollowRepository.CheckFollowByUserIdAndSellerId(userId, sellerId)
	if err != nil {
		return nil, errors.New("Erro ao verificar se usuário já segue vendedor.")
	}
	if exist {
		return nil, errors.New("Você já segue este vendedor.")
	}
	// cria a instância de UserFollow com os dados do follow
	follow := domain.UserFollow{
		FollowerID: userId,
		SellerID:   sellerId,
	}

	// cria o follow no banco de dados
	err = s.FollowRepository.CreateFollow(&follow)
	if err != nil {
		return nil, errors.New("Erro ao seguir o vendedor.")
	}

	return &follow, nil
}
