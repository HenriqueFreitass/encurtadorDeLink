package repository

import (
	"database/sql"
	"encurtador-de-link/backend/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id int) (*models.Users, error) {

	user := &models.Users{}
	err := r.db.QueryRow("Select id, name, email, password FROM users WHERE id = ?", id).Scan(
		&user.Email, &user.Password)
	return user, err

}
