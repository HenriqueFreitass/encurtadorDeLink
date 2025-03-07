package repository

import (
	"database/sql"
	"encurtador-de-link/backend/models"
	"fmt"
)

type IShortenerRepository interface {
	SaveShortenedURL(id, sitename, originalURL, newURL string, views int, useremail string) error
	GetOriginalURL(id string) (string, error)
	IncrementViews(id string) error
	GetUserLinks(userId string) ([]models.Shortener, error)
}

type ShortenerRepository struct {
	db *sql.DB
}

func NewShortenerRepository(db *sql.DB) IShortenerRepository {
	return &ShortenerRepository{db: db}
}

func (r *ShortenerRepository) SaveShortenedURL(id, sitename, originalURL, newURL string, views int, userid string) error {
	query := `
		INSERT INTO Shortener (id, sitename, originalurl, newurl, views, userid)
		VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, id, sitename, originalURL, newURL, views, userid)
	return err
}

func (r *ShortenerRepository) GetOriginalURL(id string) (string, error) {
	query := "SELECT originalurl FROM Shortener WHERE id = ?"
	var originalURL string
	err := r.db.QueryRow(query, id).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("URL not found")
		}
		return "", err
	}
	return originalURL, nil
}

func (r *ShortenerRepository) IncrementViews(id string) error {
	query := "UPDATE Shortener SET views = views + 1 WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ShortenerRepository) GetUserLinks(userId string) ([]models.Shortener, error) {
	query := "SELECT id, sitename, originalurl, newurl, views FROM Shortener WHERE userid = ?"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []models.Shortener
	for rows.Next() {
		var link models.Shortener
		if err := rows.Scan(&link.Id, &link.SiteName, &link.OriginalUrl, &link.NewUrl, &link.Views); err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return links, nil
}