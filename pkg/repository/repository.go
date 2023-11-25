package repository

import (
	"shop/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(newUser models.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
