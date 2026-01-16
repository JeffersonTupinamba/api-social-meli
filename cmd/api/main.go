package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// define o modelo de usuário para o body da requisição
type RequestUserCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"` // "seller" ou "customer"

}

type RequestUserUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// define o modelo de usuário para o banco de dados
type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null;uniqueIndex"`
	Role  string `json:"role" gorm:"not null"` // "seller" ou "customer"

}

func main() {
	r := gin.Default() // cria um novo router

	// CRIA UM NOVO USUÁRIO
	r.POST("/users", func(c *gin.Context) {
		// abre a conexão com o banco de dados
		dsn := "host=localhost user=app password=app123 dbname=social_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// cria um novo usuário com o body da requisição
		userRequest := RequestUserCreate{}
		c.ShouldBindJSON(&userRequest)
		user := User{
			Name:  userRequest.Name,
			Email: userRequest.Email,
			Role:  userRequest.Role,
		}

		result := db.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, user) // retorna o usuário criado
	})

	// BUSCA UM USUÁRIO PELO ID
	r.GET("/users/:id", func(c *gin.Context) {
		dsn := "host=localhost user=app password=app123 dbname=social_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userIDStr := c.Param("id")             // pega o id passado na url como string para converter para inteiro
		userID, err := strconv.Atoi(userIDStr) // converte o id para inteiro para buscar no banco de dados
		if err != nil {
			fmt.Println("Erro ao converter ID para inteiro:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		userfound := User{}
		db.Model(&User{}).Where("id = ?", userID).First(&userfound)
		// verifica se o usuário existe
		if userfound.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		user := User{}
		db.Model(&User{}).Where("id = ?", userID).First(&user)
		c.JSON(http.StatusOK, user) // retorna o usuário encontrado
	})

	// RETORNA TODOS OS USUÁRIOS
	r.GET("/users", func(c *gin.Context) {
		dsn := "host=localhost user=app password=app123 dbname=social_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// busca todos os usuários no banco de dados
		users := []User{}
		db.Find(&users)
		c.JSON(http.StatusOK, users)

	})
	// ATUALIZA UM USUÁRIO PELO ID
	r.PUT("/users/:id", func(c *gin.Context) {
		dsn := "host=localhost user=app password=app123 dbname=social_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userIDStr := c.Param("id")             // pega o id passado na url como string para converter para inteiro
		userID, err := strconv.Atoi(userIDStr) // converte o id para inteiro para buscar no banco de dados
		if err != nil {
			fmt.Println("Erro ao converter ID para inteiro:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		// pega o body da requisição para atualizar o usuário com o id passado na url
		userRequestUpdate := RequestUserUpdate{}
		c.ShouldBindJSON(&userRequestUpdate)

		// busca o usuário no banco de dados
		userfound := User{}
		db.Model(&User{}).Where("id = ?", userID).First(&userfound)
		// verifica se o usuário existe
		if userfound.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		user := User{
			ID:    userID,
			Name:  userRequestUpdate.Name,
			Email: userRequestUpdate.Email,
			Role:  userfound.Role, // mantém o role do usuário encontrado
		}
		// atualiza o usuário no banco de dados
		db.Model(&User{}).Where("id = ?", userID).Updates(user)
		c.JSON(http.StatusOK, user)

	})

	// DELETA UM USUÁRIO PELO ID
	r.DELETE("/users/:id", func(c *gin.Context) {
		dsn := "host=localhost user=app password=app123 dbname=social_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userIDStr := c.Param("id")             // pega o id passado na url como string para converter para inteiro
		userID, err := strconv.Atoi(userIDStr) // converte o id para inteiro para buscar no banco de dados
		if err != nil {
			fmt.Println("Erro ao converter ID para inteiro:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		// busca o usuário no banco de dados
		userfound := User{}
		db.Model(&User{}).Where("id = ?", userID).First(&userfound)
		// verifica se o usuário existe
		if userfound.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		// deleta o usuário no banco de dados
		db.Model(&User{}).Where("id = ?", userID).Delete(&userfound)
		c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
	})

	// inicia o servidor na porta 8080
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("failed to run server: %v", err) // retorna um erro se o servidor não iniciar
	}
}
