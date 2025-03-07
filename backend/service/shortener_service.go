package service

import (
	"crypto/rand"
	"encoding/base64"
	"encurtador-de-link/backend/repository"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ShortenerService interface {
	ShortenURL(originalURL, userId string) (string, error)
	GetOriginalURL(id string) (string, error)
}

type NewShortenerService struct {
	shortenerRepo repository.IShortenerRepository
}

func NewService(shortenerRepo repository.IShortenerRepository) ShortenerService {
	return &NewShortenerService{shortenerRepo: shortenerRepo}
}

func generateShortCode() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}

func extractSiteName(originalURL string) string {

	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return ""
	}
	host := parsedURL.Hostname()
	if len(host) > 4 && host[:4] == "www." {
		host = host[4:]
	}
	// Separar o nome do domínio do ".com", ".org" etc.
	parts := strings.Split(host, ".")

	// Aqui, vamos pegar o primeiro componente do domínio, que normalmente é o nome do site.
	siteName := parts[0]

	// Capitalizar a primeira letra
	// Capitaliza a primeira letra e deixa o restante minúsculo
	caser := cases.Title(language.English) // Pode ser ajustado para outras linguagens
	siteName = caser.String(siteName)
	// Retorna o nome do site formatado
	return siteName
}

func (s *NewShortenerService) ShortenURL(originalURL, userId string) (string, error) {
	shortCode := generateShortCode()

	newURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)

	sitename := extractSiteName(originalURL)

	views := 0

	err := s.shortenerRepo.SaveShortenedURL(shortCode, sitename, originalURL, newURL, views, userId)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}

func (s *NewShortenerService) GetOriginalURL(id string) (string, error) {
	originalURL, err := s.shortenerRepo.GetOriginalURL(id)
	if err != nil {
		return "", err
	}

	if !strings.Contains(originalURL, "://") {
		originalURL = "http://www." + originalURL
	}

	err = s.shortenerRepo.IncrementViews(id)
	if err != nil {
		return "", err
	}
	
	return originalURL, nil
}
