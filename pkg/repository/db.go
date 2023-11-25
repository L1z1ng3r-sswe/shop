package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Configs struct {
	Port       string
	Host       string
	User       string
	DB         string
	DBPassword string
}

func InitDB(configs Configs) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",
		configs.User, configs.DB, configs.DBPassword, configs.Host, configs.Port)
	log.Println(dsn)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
