package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"shop/configs"
	"shop/pkg/handler"
	"shop/pkg/repository"
	"shop/pkg/service"
	"shop/server"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error occured on .env: ", err.Error())
	}
	if err := configs.InitConfigs(); err != nil {
		log.Fatal("Failed to read configs: ", err.Error())
	}

	dbConfigs := repository.Configs{
		Port:       viper.GetString("database.port"),
		Host:       viper.GetString("database.host"),
		User:       viper.GetString("database.user"),
		DB:         viper.GetString("database.db"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
	db, err := repository.InitDB(dbConfigs)
	if err != nil {
		log.Fatal("Failed to connect database: ", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.SetUpRoutes()); err != nil {
			log.Fatal("Failed to run server: ", err.Error())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Failed to server shutting down: ", err.Error())
	}

	log.Println("Server gracefully stopped")

	if err := db.Close(); err != nil {
		log.Fatal("Failed to db connection close: ", err.Error())
	}
}
