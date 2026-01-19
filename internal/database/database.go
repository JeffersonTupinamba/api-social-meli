package database

import (
	"log"

	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDatabase é a função que inicializa o banco de dados e retorna uma instância do banco de dados
// e retorna um erro caso ocorra algum problema na conexão com o banco de dados
func InitDatabase() (*gorm.DB, error) {
	dns := "host=localhost user=app password=app123 dbname=social_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados:", err)
	}

	// AutoMigrate cria ou atualiza as tabelas automaticamente
	db.AutoMigrate(&domain.User{}, &domain.UserFollow{})

	return db, nil
}
