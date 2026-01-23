package handler

import (
	"net/http"
	"time"

	"github.com/JeffersonTupinamba/api-social-meli/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostHandler struct {
	DB *gorm.DB // banco de dados para realizar as operações no banco de dados (injeção de dependência) (é uma instância do banco de dados)
}

// CreatePost godoc
// @Summary      Criar uma nova postagem
// @Description  Cria uma postagem de produto para um vendedor específico
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      domain.RequestPostCreate  true  "Dados da postagem"
// @Success      200   {object}  domain.Post
// @Failure      400   {object}  map[string]string "Erro de validação"
// @Router       /products/publish [post]
func (h *PostHandler) CreatePost(c *gin.Context) {

	var req domain.RequestPostCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Validar se vendedor existe e se é vendedor
	var user domain.User
	if err := h.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vendedor não encontrado"})
		return
	}

	if user.Role != "seller" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Apenas vendedores podem criar posts"})
		return
	}

	// 2. Converter data string para time.Time
	date, err := time.Parse("02-01-2006", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido. Use DD-MM-YYYY"})
		return
	}

	// 3. Criar o post no banco de dados
	post := domain.Post{
		UserID:   req.UserID,
		Date:     date,
		Product:  req.Product,
		Category: req.Category,
		Price:    req.Price,
	}

	if err := h.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar post"})
		return
	}

	c.JSON(http.StatusOK, post)
}
