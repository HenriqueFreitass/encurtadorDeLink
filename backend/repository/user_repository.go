package repository

import (
	"database/sql"
	"encurtador-de-link/backend/models"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(Id int) (*models.Users, error) {
	user := &models.Users{}
	err := r.db.QueryRow("Select name FROM users WHERE id = ?", Id).Scan(
		&user.Name)
	return user, err

}

func (r *UserRepository) GetUserByEmail(Email string) (*models.Users, error) {

	user := &models.Users{}
	err := r.db.QueryRow("Select id, email, password FROM users WHERE email = ?", Email).Scan(
		&user.Id, &user.Email, &user.Password)
	return user, err

}

func (r *UserRepository) CreateUser(u *models.Users) (*models.Users, error) {

	query := "INSERT INTO Users(email, password, name) VALUES (?,?,?)"
	result, err := r.db.Exec(query, u.Email, u.Password, u.Name)
	if err != nil {
		return nil, errors.New("já existe um usuário com este email")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	u.Id = int(id)
	return u, nil
}
