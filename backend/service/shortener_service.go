package service

import (
	"crypto/rand"
	"encoding/base64"
	"encurtador-de-link/backend/models"
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
	GetUserLinks(userId string) ([]models.Shortener, error)
}

type NewShortenerService struct {
	shortenerRepo repository.IShortenerRepository
}

func NewService(shortenerRepo repository.IShortenerRepository) ShortenerService {
	return &NewShortenerService{shortenerRepo: shortenerRepo}
}

func (s *NewShortenerService) GetUserLinks(userId string) ([]models.Shortener, error) {
	links, err := s.shortenerRepo.GetUserLinks(userId)
	if err != nil {
		return nil, err
	}
	return links, nil
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

	parts := strings.Split(host, ".")

	siteName := parts[0]

	caser := cases.Title(language.English)
	siteName = caser.String(siteName)

	return siteName
}

func (s *NewShortenerService) ShortenURL(originalURL, userId string) (string, error) {
	shortCode := generateShortCode()

	newURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)

	if !strings.Contains(originalURL, "://") && !strings.Contains(originalURL, "www.") {
		originalURL = "http://" + originalURL
	}

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
		originalURL = "http://" + originalURL
	}

	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return "", err
	}

	originalURL = parsedURL.String()

	err = s.shortenerRepo.IncrementViews(id)
	if err != nil {
		return "", err
	}

	return originalURL, nil
}
