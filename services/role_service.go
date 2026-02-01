package services

import (
	"project-golang/models"
	"project-golang/repositories"
)

type RoleService struct {
	repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) GetAll() ([]models.Role, error) {
	return s.repo.GetAll()
}