package database

import (
	"fmt"
	"log"
	"os"

	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase é a função que inicializa o banco de dados e retorna uma instância do banco de dados
// e retorna um erro caso ocorra algum problema na conexão com o banco de dados
func ConnectDatabase() (*gorm.DB, error) {
	// Pega as configurações do banco de dados das variáveis de ambiente
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	// Monta a string de conexão (DSN)
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados:", err)
	}

	// AutoMigrate cria ou atualiza as tabelas automaticament

	return db, db.AutoMigrate(&domain.User{}, &domain.UserFollow{}, &domain.Post{})

}
