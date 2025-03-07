package handlers

import (
	"encurtador-de-link/backend/service"

	"github.com/gin-gonic/gin"
)

type ShortenerHandler struct {
	shortenerservice service.ShortenerService
}

func NewShortenerHandler(shortenerservice service.ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{shortenerservice: shortenerservice}
}

func (h *ShortenerHandler) ShortenURL(c *gin.Context) {
	var body struct {
		OriginalURL string `json:"url" binding:"required"`
		UserId      string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	shortCode, err := h.shortenerservice.ShortenURL(body.OriginalURL, body.UserId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error saving URL"})
		return
	}

	c.JSON(200, gin.H{"short_url": "http://localhost:8080/" + shortCode})
}

func (h *ShortenerHandler) RedirectURL(c *gin.Context) {
	shortCode := c.Param("id")

	originalURL, err := h.shortenerservice.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(302, originalURL)
}

func (h *ShortenerHandler) GetUserLinks(c *gin.Context) {
	userId := c.Param("id")

	// Chama o serviço para obter os links do usuário
	links, err := h.shortenerservice.GetUserLinks(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching links"})
		return
	}

	// Responde com os links do usuário
	c.JSON(200, gin.H{"links": links})
}
