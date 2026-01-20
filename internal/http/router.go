package http

import (
	"github.com/JeffersonTupinamba/api-social-meli/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// setupRouter é a função que configura o router
func SetupRouter(db *gorm.DB) *gin.Engine {

	r := gin.Default()                // cria uma nova instância de gin default que é um router padrão do gin
	h := &handler.UserHandler{DB: db} // cria uma nova instância de UserHandler com o banco de dados

	// US 0013: Criar um novo usuário CRUD (Implícito)
	r.POST("/users", h.CreateUser)
	r.GET("/users/:userId", h.GetUser)
	r.GET("/users", h.GetUsers)
	r.PUT("/users/:userId", h.UpdateUser)
	r.DELETE("/users/:userId", h.DeleteUser)

	// US 0001: Poder "seguir" um vendedor específico
	r.POST("/users/:userId/follow/:sellerId", h.FollowUser)

	// US 0002: Obter o número de seguidores de um vendedor
	r.GET("/users/:userId/followers/count", h.GetUsersFollowersCountBySeller)

	// US 0003: Obter uma lista de todos os usuários que seguem um determinado vendedor (Quem me segue?)
	r.GET("/users/:userId/followers/list", h.GetUsersFollowersList)

	// US 0004: Obter uma lista de todos os vendedores seguidos por um determinado usuário (Quem estou seguindo?)
	r.GET("/users/:userId/followed/list", h.GetUsersFollowedSellersList)

	// US 0007: Para que você possa "Unfollow" um determinado vendedor.
	r.DELETE("/users/:userId/unfollow/:sellerId", h.UnfollowUser)

	return r // retorna o router configurado
}
