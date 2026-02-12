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
//	@Accept			json
//	@Produce		json
//	@Param			userId			path		int					true	"ID do usuário que vai seguir"
//	@Param			sellerId	path		int					true	"ID do vendedor a ser seguido"
//	@Success		200				{string}	string				"Usuário seguido com sucesso"
//	@Failure		400				{object}	map[string]string	"Erro na requisição"
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Você não pode seguir você mesmo."})
			return
		case "Vendedor não encontrado.":
			c.JSON(http.StatusNotFound, gin.H{"error": "Vendedor não encontrado."})
			return
		case "O ID informado não pertence a um vendedor.":
			c.JSON(http.StatusBadRequest, gin.H{"error": "O ID informado não pertence a um vendedor."})
			return
		case "Você já segue este vendedor.":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Você já segue este vendedor."})
			return
		case "Usuário não encontrado.":
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário seguindo vendedor com sucesso!",
		"data": follow})
}
