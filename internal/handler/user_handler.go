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

// CRIA UM NOVO USUÁRIO

// CreateUser godoc
// @Summary      Criar um novo usuário
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      domain.RequestUserCreate  true  "Dados do usuário"
// @Success      201   {object}  domain.User
// @Failure      400   {object}  map[string]string "Dados inválidos ou JSON malformado"
// @Failure      500   {object}  map[string]string "Erro interno ao salvar no banco"
// @Router       /users [post]
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário."})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// BUSCA UM USUÁRIO PELO ID

// GetUser godoc
// @Summary      Buscar usuário por ID
// @Tags         users
// @Produce      json
// @Param        userId  path      int  true  "ID do Usuário"
// @Success      200     {object}  domain.User
// @Failure      400     {object}  map[string]string "ID inválido (não é um número)"
// @Failure      404     {object}  map[string]string "Usuário não encontrado no banco"
// @Router       /users/{userId} [get]
func (h *UserHandler) GetUser(c *gin.Context) {

	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido."}) //status bad request é um status que indica que a requisição é inválida
		return
	}

	var user domain.User
	if err := h.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."}) //status not found é um status que indica que o recurso não foi encontrado
		return
	}

	c.JSON(http.StatusOK, user)
}

// RETORNA TODOS OS USUÁRIOS

// GetUsers godoc
// @Summary      Listar todos os usuários
// @Description  Retorna uma lista contendo todos os usuários cadastrados no sistema
// @Tags         users
// @Produce      json
// @Success      200  {array}   domain.User        "Lista de usuários retornada com sucesso"
// @Failure      500  {object}  map[string]string  "Erro interno ao buscar usuários no banco"
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {

	var users []domain.User

	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários."})
		return
	}

	c.JSON(http.StatusOK, users)
}

// ATUALIZA UM USUÁRIO PELO ID

// UpdateUser godoc
// @Summary      Atualizar um usuário
// @Description  Atualiza os dados (nome e email) de um usuário existente pelo seu ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        userId  path      int                        true  "ID do usuário a ser atualizado"
// @Param        user    body      domain.RequestUserUpdate   true  "Novos dados do usuário"
// @Success      200     {object}  domain.User                "Usuário atualizado com sucesso"
// @Failure      400     {object}  map[string]string          "ID inválido ou JSON malformado"
// @Failure      404     {object}  map[string]string          "Usuário não encontrado"
// @Router       /users/{userId} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {

	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido."})
		return
	}

	var userfound domain.User
	if err := h.DB.First(&userfound, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
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
		Role:  userRequestUpdate.Role,
	})

	c.JSON(http.StatusOK, userfound)
}

// DELETA UM USUÁRIO PELO ID

// DeleteUser godoc
// @Summary      Deletar um usuário
// @Tags         users
// @Produce      json
// @Param        userId  path      int  true  "ID do Usuário"
// @Success      200     {object}  map[string]string "Mensagem de sucesso"
// @Failure      400     {object}  map[string]string "ID inválido"
// @Failure      404     {object}  map[string]string "Usuário não existe para ser deletado"
// @Router       /users/{userId} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido."})
		return
	}

	result := h.DB.Delete(&domain.User{}, userId)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso."})
}

// US 0001: Poder "seguir" um vendedor específico

// FollowUser godoc
// @Summary      Seguir um vendedor
// @Description  Permite que um usuário siga um vendedor específico
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        userId       path      int  true  "ID do usuário que vai seguir"
// @Param        userIdToFollow  path      int  true  "ID do vendedor a ser seguido"
// @Success      200          {string}  string "Usuário seguido com sucesso"
// @Failure      400          {object}  map[string]string "Erro na requisição"
// @Router       /users/{userId}/follow/{userIdToFollow} [post]
func (h *UserHandler) FollowUser(c *gin.Context) {

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

	// cria o follow no banco de dados
	if err := h.DB.Create(&follow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar follow: "})
		return
	}

	// retorna a mensagem de sucesso e os dados do follow
	c.JSON(http.StatusOK, gin.H{"message": "Usuário seguindo vendedor com sucesso!.",
		"data": follow})
}

// US 0002: Obter o número de seguidores de um vendedor
// GetUsersFollowersCountBySeller godoc
// @Summary      Obter contagem de seguidores
// @Description  Retorna a quantidade total de seguidores de um usuário, DESDE QUE ele seja um vendedor (role='seller')
// @Tags         users
// @Param        userId  path  int  true  "ID do vendedor"
// @Failure      400     {object}  map[string]string "ID inválido ou usuário não é vendedor"
// @Router       /users/{userId}/followers/count [get]
func (h *UserHandler) GetUsersFollowersCountBySeller(c *gin.Context) {
	userIdStr := c.Param("userId") // Alinhado com o roteador
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido."})
		return
	}
	// 1. Verificar se o vendedor existe e se a role é válida
	var user domain.User
	if err := h.DB.Where("id = ? AND role = ?", userId, "seller").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vendedor não encontrado ou usuário não é um vendedor."})
		return
	}

	// 2. Contar os seguidores
	var count int64
	if err := h.DB.Model(&domain.UserFollow{}).Where("seller_id = ?", userId).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao contar seguidores."})
		return
	}

	response := domain.UserFollowersCountResponse{
		UserID:         user.ID,
		UserName:       user.Name,
		FollowersCount: count,
	}

	c.JSON(http.StatusOK, response)
}

// US 0003: Obter lista de seguidores de um vendedor

// GetFollowersList godoc
// @Summary      Listar seguidores de um vendedor
// @Description  Retorna a lista de todos os usuários que seguem um vendedor específico
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        userId  path      int  true  "ID do vendedor"
// @Success      200     {object}  domain.UserFollowersListResponse
// @Failure      404     {object}  map[string]string "Vendedor não encontrado"
// @Router       /users/{userId}/followers/list [get]
func (h *UserHandler) GetUsersFollowersList(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	// 1. Verificar se o vendedor existe e se é um vendedor
	var seller domain.User
	if err := h.DB.Where("id = ? AND role = ?", userId, "seller").First(&seller).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vendedor não encontrado ou usuário não é um vendedor"})
		return
	}

	// 2. Buscar os seguidores (fazendo um Join com a tabela de usuários)
	var followers []domain.FollowerDTO
	err := h.DB.Table("users").
		Select("users.id as user_id, users.name as user_name").
		Joins("INNER JOIN user_follows ON user_follows.follower_id = users.id").
		Where("user_follows.seller_id = ?", userId).
		Scan(&followers).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar seguidores"})
		return
	}

	// 3. Montar a resposta final
	response := domain.UserFollowersListResponse{
		UserID:    seller.ID,
		UserName:  seller.Name,
		Followers: followers,
	}

	c.JSON(http.StatusOK, response)
}

// US 0004: Obter uma lista de todos os vendedores seguidos por um determinado usuário (Quem estou seguindo?)

// GetFollowersList godoc
// @Summary      Listar todos os vendedores seguidos
// @Description  Retorna a lista de todos os vendedores seguidos por um determinado usuário
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        userId  path      int  true  "ID do vendedor"
// @Success      200     {object}  domain.UserFollowersListResponse
// @Failure      404     {object}  map[string]string "Vendedor não encontrado"
// @Router       /users/{userId}/followed/list [get]
func (h *UserHandler) GetUsersFollowedSellersList(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	// 1. Verificar se o usuário existe
	var user domain.User
	if err := h.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// 2. Buscar os vendedores seguidos (fazendo um Join com a tabela de usuários)
	var followedSellers []domain.FollowerDTO
	err := h.DB.Table("users").
		Select("users.id as user_id, users.name as user_name").
		Joins("INNER JOIN user_follows ON user_follows.seller_id = users.id").
		Where("user_follows.follower_id = ?", userId).
		Scan(&followedSellers).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar vendedores seguidos"})
		return
	}

	// 3. Montar a resposta final
	response := domain.UserFollowedListResponse{
		UserID:          user.ID,
		UserName:        user.Name,
		FollowedSellers: followedSellers,
	}

	c.JSON(http.StatusOK, response)
}

// US 0007: Para que você possa "Unfollow" um determinado vendedor.

// UnfollowUser godoc
// @Summary      Deixar de seguir um vendedor
// @Description  Permite que um usuário pare de seguir um vendedor específico
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        userId         path      int  true  "ID do usuário"
// @Param        userIdToUnfollow  path      int  true  "ID do vendedor a deixar de seguir"
// @Success      200            {string}  string "Deixou de seguir com sucesso"
// @Router       /users/{userId}/unfollow/{userIdToUnfollow} [put]
func (h *UserHandler) UnfollowUser(c *gin.Context) {

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do vendedor inválido."})
		return
	}

	// Verifica se o vendedor existe e se é um vendedor
	var seller domain.User
	if err := h.DB.Where("id = ? AND role = ?", sellerId, "seller").First(&seller).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vendedor não encontrado ou usuário não é um vendedor."})
		return
	}

	// deleta o follow no banco de dados e retorna um erro caso ocorra algum problema
	result := h.DB.Where("follower_id = ? AND seller_id = ?", userId, sellerId).Delete(&domain.UserFollow{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar follow."})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Você não segue este vendedor."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deixou de seguir vendedor com sucesso!"})
}
