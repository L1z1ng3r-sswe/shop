package service

import (
	"os"
	"shop/models"
	"shop/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repos repository.Authorization
}

func NewAuthService(r repository.Authorization) *authService {
	return &authService{repos: r}
}

func (s *authService) CreateUser(newUser models.User) (int, error) {
	newUser.Password = hashedPassword(newUser.Password)

	id, err := s.repos.CreateUser(newUser)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func hashedPassword(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password+os.Getenv("HASH_SALT")), 10)
	return string(passwordHash)
}
