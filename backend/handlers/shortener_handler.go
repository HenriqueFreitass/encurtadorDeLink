package handlers

import (
	"encurtador-de-link/backend/service"

	"github.com/gin-gonic/gin"
)

// Handler lida com as requisições HTTP
type ShortenerHandler struct {
	service service.NewShortenerService
}

// NewHandler cria uma nova instância do handler
func NewShortenerHandler(s service.NewShortenerService) *ShortenerHandler {
	return &ShortenerHandler{service: s}
}

// ShortenURL lida com a requisição para encurtar uma URL
func (h *ShortenerHandler) ShortenURL(c *gin.Context) {
	var body struct {
		OriginalURL string `json:"originalurl" binding:"required"`
		UserEmail   string `json:"useremail" binding:"required"`
	}

	// Parse the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Chama o serviço para encurtar a URL
	shortCode, err := h.service.ShortenURL(body.OriginalURL, body.UserEmail)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error saving URL"})
		return
	}

	// Retorna a URL encurtada com o código curto gerado
	c.JSON(200, gin.H{"short_url": "http://localhost:5173/" + shortCode})
}

// RedirectURL lida com a requisição de redirecionamento de um código curto
func (h *ShortenerHandler) RedirectURL(c *gin.Context) {
	shortCode := c.Param("id")

	// Chama o serviço para recuperar a URL original
	originalURL, err := h.service.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(302, originalURL)
}
