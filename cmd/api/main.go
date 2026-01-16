package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// RequestUserCreate define o que esperamos receber no POST /users
type RequestUserCreate struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required"`
}

// RequestUserUpdate define o que permitimos atualizar no PUT /users/:id
type RequestUserUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
}

// User é a representação da nossa tabela no banco de dados
type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null;uniqueIndex"`
	Role  string `json:"role" gorm:"not null"`
}

func main() {
	// 1. CONEXÃO ÚNICA: Abrimos o banco uma vez no início do programa
	dsn := "host=localhost user=app password=app123 dbname=social_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados:", err)
	}

	// AutoMigrate cria ou atualiza as tabelas automaticamente
	db.AutoMigrate(&User{})

	r := gin.Default()

	// --- ROTAS ---

	// CRIA UM NOVO USUÁRIO
	r.POST("/users", func(c *gin.Context) {
		var userRequest RequestUserCreate

		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
			return
		}

		user := User{
			Name:  userRequest.Name,
			Email: userRequest.Email,
			Role:  userRequest.Role,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	// BUSCA UM USUÁRIO PELO ID
	r.GET("/users/:id", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var user User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	// RETORNA TODOS OS USUÁRIOS
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	// ATUALIZA UM USUÁRIO PELO ID
	r.PUT("/users/:id", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var userfound User
		if err := db.First(&userfound, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}

		var userRequestUpdate RequestUserUpdate
		if err := c.ShouldBindJSON(&userRequestUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Model(&userfound).Updates(User{
			Name:  userRequestUpdate.Name,
			Email: userRequestUpdate.Email,
		})

		c.JSON(http.StatusOK, userfound)
	})

	// DELETA UM USUÁRIO PELO ID
	r.DELETE("/users/:id", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		result := db.Delete(&User{}, userID)

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
	})

	fmt.Println("Servidor rodando na porta 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Falha ao iniciar o servidor:", err)
	}
}
