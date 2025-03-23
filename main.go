package main

import (
	"auth-service/config"
	"auth-service/internal/handlers"
	"auth-service/internal/repositories"
	"auth-service/internal/services"
	"auth-service/routes"
	"log"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repositories
	authRepo := repositories.NewAuthRepository(db)

	// Initialize services
	authService := services.NewAuthService(authRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Setup router
	router := routes.SetupRouter(authHandler)

	// Start server
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
