package service

import (
	"encurtador-de-link/backend/models"
	"encurtador-de-link/backend/repository"
	"errors"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserProfile(userID int) (*models.Users, error) {
	return s.userRepo.GetUserByID(userID)
}

func (s *UserService) AuthenticateUser(email, password string) (*models.Users, error) {
	// Buscar usuário no banco pelo email
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Comparar senhas (o ideal seria usar hash)
	if user.Password != password {
		return nil, errors.New("senha incorreta")
	}

	return user, nil
}

func (s *UserService) CreateUser(user *models.Users) (*models.Users, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("algum dos campos não foi preenchido")
	}
	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
