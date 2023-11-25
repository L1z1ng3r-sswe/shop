package models

import "time"

type User struct {
	ID       int       `json:"-" db:"id" validate:"-"`
	Username string    `db:"username" validate:"required,min=3,max=128"`
	Lastname string    `db:"lastname" validate:"required,min=3,max=128"`
	Email    string    `db:"email" validate:"required,email,max=128"`
	Password string    `db:"password" validate:"required,min=6,max=256"`
	Birthday time.Time `db:"birthday" validate:"required"`
}
