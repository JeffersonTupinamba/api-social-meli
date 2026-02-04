package handler

import (
	"net/http"
	"strconv"

	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"github.com/JeffersonTupinamba/api-social-meli/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FollowHandler é o handler para as operações de follow (injeção de dependência)
type FollowHandler struct {
	DB            *gorm.DB
	FollowService *service.FollowService
}

// US 0001: Poder "seguir" um vendedor específico

// FollowUser godoc
//
//	@Summary		Seguir um vendedor
//	@Description	Permite que um usuário siga um vendedor específico
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			userId			path		int					true	"ID do usuário que vai seguir"
//	@Param			userIdToFollow	path		int					true	"ID do vendedor a ser seguido"
//	@Success		200				{string}	string				"Usuário seguido com sucesso"
//	@Failure		400				{object}	map[string]string	"Erro na requisição"
//	@Router			/users/{userId}/follow/{userIdToFollow} [post]
func (h *FollowHandler) FollowUser(c *gin.Context) {

	userIdStr := c.Param("userId")
	sellerIdStr := c.Param("sellerId")
	// converte o ID do usuário para string
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido."})
		return
	}
	sellerId, err := strconv.Atoi(sellerIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID é do vendedor inválido."})
		return
	}
	// verifica se o usuário está tentando seguir ele mesmo
	if userId == sellerId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Você não pode seguir você mesmo."})
		return
	}
	// verifica se o vendedor existe
	var existingSeller domain.User
	if err := h.DB.First(&existingSeller, sellerId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vendedor não encontrado."})
		return
	}
	if existingSeller.Role != "seller" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O ID informado não pertence a um vendedor."})
		return
	}
	// verifica se o usuário existe
	var existingUser domain.User
	if err := h.DB.First(&existingUser, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
		return
	}
	// verifica se o usuário já está seguindo o vendedor
	var existingFollow domain.UserFollow
	if err := h.DB.Where("follower_id = ? AND seller_id = ?", userId, sellerId).First(&existingFollow).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Você já segue este vendedor."})
		return
	}
	// cria a instância de UserFollow com os dados do follow
	follow := domain.UserFollow{
		FollowerID: userId,
		SellerID:   sellerId,
	}

	// retorna a mensagem de sucesso e os dados do follow
	c.JSON(http.StatusOK, gin.H{"message": "Usuário seguindo vendedor com sucesso!.",
		"data": follow})
}
