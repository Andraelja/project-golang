package services

import (
	"errors"
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
	role, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(role) < 1 {
		return nil, errors.New("Data is still empty, please add data!")
	}

	return role, nil
}

func (s *RoleService) Create(data *models.Role) error {
	return s.repo.Create(data)
}

func (s *RoleService) GetByID(id int) (*models.Role, error) {
	role, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return nil, errors.New("Data not found!")
	}
	
	return role, nil
}
