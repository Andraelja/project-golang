package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) Create(data *models.User) error {
	if data.RoleID == 0 {
		return errors.New("role cannot be empty")
	}

	role, err := s.roleRepo.GetByID(data.RoleID)
	if err != nil {
		return err
	}
	if role == nil {
		return errors.New("role not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	data.Password = string(hashedPassword)

	return s.userRepo.Create(data)
}

func (s *UserService) GetByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found!")
	}

	return user, nil
}

func (s *UserService) Update(user *models.User) error {
	if user.RoleID == 0 {
		return errors.New("Role cannot empty!")
	}

	role, err := s.roleRepo.GetByID(user.RoleID)
	
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if role == nil {
		return errors.New("role not found!")
	}

	return s.userRepo.Update(user)
}
