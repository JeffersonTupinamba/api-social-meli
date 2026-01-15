package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// define o modelo de usuário para o body da requisição
type RequestUserCreate struct {
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;uniqueIndex"`
	Role  string `gorm:"not null"` // "seller" ou "customer"

}

// define o modelo de usuário para o banco de dados
type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;uniqueIndex"`
	Role  string `gorm:"not null"` // "seller" ou "customer"

}

func main() {
	r := gin.Default() // cria um novo router

	// cria um novo usuário
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

	// retorna todos os usuários
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
	// inicia o servidor na porta 8080
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("failed to run server: %v", err) // retorna um erro se o servidor não iniciar
	}
}
