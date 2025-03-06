package service

import (
	"crypto/rand"
	"encoding/base64"
	"encurtador-de-link/backend/repository"
	"fmt"
	"net/url"
)

// Service é a interface que define os métodos de negócio
type ShortenerService interface {
	ShortenURL(originalURL, useremail string) (string, error)
	GetOriginalURL(id string) (string, error)
}

type NewShortenerService struct {
	shortenerRepo repository.NewShortenerRepositoryy
}

// NewService cria uma nova instância do serviço de encurtamento de URL
func NewService(shortenerRepo repository.NewShortenerRepositoryy) ShortenerService {
	return &NewShortenerService{shortenerRepo: shortenerRepo}
}

// generateShortCode gera um código curto aleatório
func generateShortCode() string {
	b := make([]byte, 6) // Tamanho do código curto
	rand.Read(b)         // Preenche com bytes aleatórios
	return base64.URLEncoding.EncodeToString(b)[:6] // Retorna os primeiros 6 caracteres
}

func extractSiteName(originalURL string) string {
	// Faz a análise da URL
	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return ""
	}

	// Retorna o nome do domínio sem o "www" ou protocolo
	host := parsedURL.Hostname()
	// Se o host começar com "www.", removemos "www."
	if len(host) > 4 && host[:4] == "www." {
		host = host[4:]
	}
	return host
}

// ShortenURL gera um código curto para a URL original e a armazena no banco de dados
func (s *NewShortenerService) ShortenURL(originalURL, useremail string) (string, error) {
	// Gera o código curto
	shortCode := generateShortCode()

	// Gera a URL encurtada com base no código curto
	newURL := fmt.Sprintf("http://localhost:5173/%s", shortCode)

	sitename := extractSiteName(originalURL)

	// Inicializa a contagem de views como 0
	views := 0

	// Salva no banco de dados usando o código curto como o atributo id
	err := s.shortenerRepo.SaveShortenedURL(shortCode, sitename, originalURL, newURL, views, useremail)
	if err != nil {
		return "", err
	}

	// Retorna o código curto gerado
	return shortCode, nil
}

// GetOriginalURL busca a URL original a partir do código curto (id)
func (s *NewShortenerService) GetOriginalURL(id string) (string, error) {
	err := s.shortenerRepo.IncrementViews(id)
	if err != nil {
		return "", err
	}
	originalURL, err := s.shortenerRepo.GetOriginalURL(id)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
