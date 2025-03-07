package repository

import (
	"database/sql"
	"fmt"
)

// Repository interface define os métodos de acesso a dados
type IShortenerRepository interface {
	SaveShortenedURL(id, sitename, originalURL, newURL string, views int, useremail string) error
	GetOriginalURL(id string) (string, error)
	IncrementViews(id string) error // Função para incrementar as visualizações
}

type ShortenerRepository struct {
	db *sql.DB
}

// NewRepository cria uma nova instância do repositório MySQL
func NewShortenerRepository(db *sql.DB) IShortenerRepository {
	return &ShortenerRepository{db: db}
}

// SaveShortenedURL armazena o código curto (id) e os outros dados no banco de dados
func (r *ShortenerRepository) SaveShortenedURL(id, sitename, originalURL, newURL string, views int, useremail string) error {
	// Insere a URL e outros dados no banco de dados, usando o id como chave primária
	query := `
		INSERT INTO Shortener (id, sitename, originalurl, newurl, views, useremail)
		VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, id, sitename, originalURL, newURL, views, useremail)
	return err
}

// GetOriginalURL busca a URL original a partir do código curto (id)
func (r *ShortenerRepository) GetOriginalURL(id string) (string, error) {
	// Busca a URL original com base no código curto (id)
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
	// Atualiza o campo views para incrementar o contador
	query := "UPDATE Shortener SET views = views + 1 WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
