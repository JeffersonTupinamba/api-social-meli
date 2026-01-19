package handler

import (
	"net/http"
	"strconv"

	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler é o handler para as operações de usuário (injeção de dependência)
type UserHandler struct {
	DB *gorm.DB // banco de dados para realizar as operações no banco de dados (injeção de dependência) (é uma instância do banco de dados)
}

// CRIA UM NOVO USUÁRIO PELO BODY DA REQUISIÇÃO
func (h *UserHandler) CreateUser(c *gin.Context) {

	var createUser domain.RequestUserCreate // cria uma nova instância de RequestUserCreate

	if err := c.ShouldBindJSON(&createUser); err != nil { // should bind json é um middleware que valida os dados recebidos no corpo da requisição
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	user := domain.User{
		Name:  createUser.Name,
		Email: createUser.Email,
		Role:  createUser.Role,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// BUSCA UM USUÁRIO PELO ID
func (h *UserHandler) GetUser(c *gin.Context) {

	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"}) //status bad request é um status que indica que a requisição é inválida
		return
	}

	var user domain.User
	if err := h.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"}) //status not found é um status que indica que o recurso não foi encontrado
		return
	}

	c.JSON(http.StatusOK, user)
}

// RETORNA TODOS OS USUÁRIOS
func (h *UserHandler) GetUsers(c *gin.Context) {

	var users []domain.User

	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {

	// ATUALIZA UM USUÁRIO PELO ID
	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}

	var userfound domain.User
	if err := h.DB.First(&userfound, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var userRequestUpdate domain.RequestUserUpdate
	if err := c.ShouldBindJSON(&userRequestUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&userfound).Updates(domain.User{
		Name:  userRequestUpdate.Name,
		Email: userRequestUpdate.Email,
	})

	c.JSON(http.StatusOK, userfound)
}

// DELETA UM USUÁRIO PELO ID
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}

	result := h.DB.Delete(&domain.User{}, userId)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}

// US 0001: Poder "seguir" um vendedor específico
func (h *UserHandler) FollowUser(c *gin.Context) {

	userIdStr := c.Param("userId")
	sellerIdStr := c.Param("sellerId")

	// converte o ID do usuário para string
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}
	sellerId, err := strconv.Atoi(sellerIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID é do vendedor inválido"})
		return
	}

	// cria a instância de UserFollow com os dados do follow
	follow := domain.UserFollow{
		FollowerID: userId,
		SellerID:   sellerId,
	}

	// cria o follow no banco de dados
	if err := h.DB.Create(&follow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar follow: " + err.Error()})
		return
	}

	// retorna a mensagem de sucesso e os dados do follow
	c.JSON(http.StatusOK, gin.H{"message": "Usuário seguindo vendedor com sucesso!",
		"data": follow,
	})
}
