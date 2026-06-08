package main

import (
	"context"
	"log"
	"net/http"

	"local-notice-hex-go/configs"
	"local-notice-hex-go/internal/infrastructure/database/postgres"
	"local-notice-hex-go/internal/infrastructure/http/handler"
	"local-notice-hex-go/internal/infrastructure/http/router"
	"local-notice-hex-go/internal/service/user"
	"local-notice-hex-go/pkg/database"
	"local-notice-hex-go/pkg/migration"
)

func main() {
	config := configs.LoadConfig()

	dbConnector := database.NewSqlConnector()
	db, err := dbConnector.Connect(context.Background(), config.DatabaseURL)
	if err != nil {

		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	migrator := migration.NewMigrator(db)
	if err := migrator.Up(context.Background()); err != nil {

		log.Fatal("Failed to run migrations:", err)
	}

	userRepo := postgres.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, config.JWTSecret)

	r := router.SetupRouter(userHandler, config.JWTSecret)

	log.Printf("Server starting on port %s", config.ServerPort)
	if err := http.ListenAndServe(":"+config.ServerPort, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
