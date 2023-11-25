package service

import (
	"shop/models"
	"shop/pkg/repository"
)

type Authorization interface {
	CreateUser(newUser models.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
