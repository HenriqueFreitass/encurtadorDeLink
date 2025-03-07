package service

import (
	"crypto/rand"
	"encoding/base64"
	"encurtador-de-link/backend/repository"
	"fmt"
	"net/url"
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
	return host
}

func (s *NewShortenerService) ShortenURL(originalURL, userId string) (string, error) {
	shortCode := generateShortCode()

	newURL := fmt.Sprintf("http://localhost:5173/%s", shortCode)

	sitename := extractSiteName(originalURL)

	views := 0

	err := s.shortenerRepo.SaveShortenedURL(shortCode, sitename, originalURL, newURL, views, userId)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}

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
