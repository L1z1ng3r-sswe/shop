package repository

import (
	"errors"
	"shop/models"

	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (repos *authRepository) CreateUser(newUser models.User) (int, error) {
	var userId int

	if err := repos.db.QueryRow("INSERT INTO users ( username ,lastname ,email ,password ,birthday) values ($1, $2, $3, $4, $5) RETURNING id",
		newUser.Username, newUser.Lastname, newUser.Email, newUser.Password, newUser.Birthday).Scan(&userId); err != nil {
		return 0, errors.New("User is already exist: " + err.Error())
	}

	return userId, nil
}
