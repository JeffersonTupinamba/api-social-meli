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

	r.POST("/users", h.CreateUser)           // cria uma nova rota POST /users que chama a função createUser do handler
	r.GET("/users/:userId", h.GetUser)       // cria uma nova rota GET /users/:userId que chama a função getUser do handler
	r.GET("/users", h.GetUsers)              // cria uma nova rota GET /users que chama a função getUsers do handler
	r.PUT("/users/:userId", h.UpdateUser)    // cria uma nova rota PUT /users/:userId que chama a função updateUser do handler
	r.DELETE("/users/:userId", h.DeleteUser) // cria uma nova rota DELETE /users/:userId que chama a função deleteUser do handler

	r.POST("/users/:userId/follow/:sellerId", h.FollowUser) // US 0001: Poder "seguir" um vendedor específico
	// r.DELETE("/users/:userId/follow/:sellerId", h.UnfollowUser) // US 0002: Poder "desseguir" um vendedor específico
	return r // retorna o router configurado
}
