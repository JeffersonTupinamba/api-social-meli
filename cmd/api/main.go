package main

import (
	"fmt"
	"log"

	"github.com/JeffersonTupinamba/api-social-meli/internal/database"
	"github.com/JeffersonTupinamba/api-social-meli/internal/http"
	"github.com/joho/godotenv"
)

// @title			Social Meli API
// @version		1.0
// @description	API para o desafio Social Meli do bootcamp.
// @host			localhost:8080
// @BasePath		/
func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// inicializa o banco de dados
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados:", err)
	}

	// Chama a função SetupRouter para configurar as rotas do pacote http
	// passa o 'db' para que o 'router' possa entregar a requisição para o 'handlers'
	r := http.SetupRouter(db)

	// INICIA O SERVIDOR HTTP NA PORTA 8080 (MELI_PORT) DEVE SER CONFIGURADA NO ARQUIVO .env
	fmt.Println("Servidor rodando na porta 8080...")

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Falha ao iniciar o servidor:", err) //log.Fatal é uma função que termina o programa e imprime a mensagem de erro

	}

}
