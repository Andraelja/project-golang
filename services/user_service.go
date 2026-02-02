package services

import (
	"project-golang/models"
	"project-golang/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
	roleRepo *repositories.RoleRepository
}

func NewUserService(
	userRepo *repositories.UserRepository,
	roleRepo *repositories.RoleRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.userRepo.GetAll()
}
