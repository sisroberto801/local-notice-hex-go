package main

import (
	"log"
	"net/http"

	"local-notice-hex-go/configs"
	"local-notice-hex-go/internal/infrastructure/database/postgres"
	"local-notice-hex-go/internal/infrastructure/http/handler"
	"local-notice-hex-go/internal/infrastructure/http/router"
	"local-notice-hex-go/internal/service/user"
)

func main() {
	config := configs.LoadConfig()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)

	userService := user.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRouter(userHandler)

	log.Printf("Server starting on port %s", config.ServerPort)
	if err := http.ListenAndServe(":"+config.ServerPort, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}