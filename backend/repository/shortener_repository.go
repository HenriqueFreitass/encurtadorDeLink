package repository

import (
	"database/sql"
	"fmt"
)

type IShortenerRepository interface {
	SaveShortenedURL(id, sitename, originalURL, newURL string, views int, useremail string) error
	GetOriginalURL(id string) (string, error)
	IncrementViews(id string) error
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
