package handler

import (
	"net/http"
	"strconv"

	"github.com/JeffersonTupinamba/api-social-meli/internal/service"
	"github.com/gin-gonic/gin"
)

// FollowHandler é o handler para as operações de follow (injeção de dependência)
type FollowHandler struct {
	FollowService *service.FollowService
}

// US 0001: Poder "seguir" um vendedor específico

// FollowUser godoc
//
//	@Summary		Seguir um vendedor
//	@Description	Permite que um usuário(customer) siga um vendedor(seller) específico
//	@Tags			Users
//	@Produce		json
//	@Param			userId			path		int					true	"ID do usuário que vai seguir"
//	@Param			sellerId	path		int					true	"ID do vendedor a ser seguido"
//	@Success		200				{object}	map[string]interface{}	"Retorna message e data (follow)"
//	@Failure		400				{object}	map[string]string		"Erro na requisição"
//	@Failure		403				{object}	map[string]string		"Regra de roles (seller/customer)"
//	@Failure		404				{object}	map[string]string		"Usuário ou vendedor não encontrado"
//	@Failure		409				{object}	map[string]string		"Você já segue este vendedor."
//	@Failure		500				{object}	map[string]string		"Erro interno"
//	@Router			/users/{userId}/follow/{sellerId} [post]
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

	follow, err := h.FollowService.FollowUserService(userId, sellerId)
	if err != nil {
		switch err.Error() {
		case "Você não pode seguir você mesmo.":
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case "Você já segue este vendedor.":
			// Ideal: 409 Conflict (recurso já existe)
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case "Vendedor não encontrado.":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "Usuário não encontrado.":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "O ID informado não pertence a um vendedor.":
			// Ideal: 403 Forbidden (role não permite) ou 400 (input inválido)
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case "O ID informado não pertence a um comprador.":
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case "Erro ao verificar se usuário já segue vendedor.":
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		case "Erro ao seguir o vendedor.":
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro inesperado ao seguir vendedor."})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário seguindo vendedor com sucesso!",
		"data":    follow,
	})
}
