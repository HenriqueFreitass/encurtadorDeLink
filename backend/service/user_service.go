package service

import (
	"encurtador-de-link/backend/models"
	"encurtador-de-link/backend/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserRepository(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserProfile(userID int) (*models.Users, error) {
	return s.userRepo.GetUserByID(userID)
}
