package services

import (
	"errors"
	"log"
	"project-golang/repositories"
	"project-golang/utils"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", err
	}

	if user == nil {
		log.Printf("User not found for username: %s", username)
		return "", errors.New("invalid username or password")
	}

	// cek password
	log.Printf("Login attempt for username: %s, received password: %s", username, password)
	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		log.Printf("Password check failed for username: %s, error: %v", username, err)
		return "", errors.New("invalid username or password")
	}

	// generate JWT
	token, err := utils.GenerateToken(user.ID, user.Role.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}
