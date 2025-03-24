package main

import (
	"auth-service/config"
	"auth-service/internal/handlers"
	"auth-service/internal/repositories"
	"auth-service/internal/services"
	"auth-service/routes"
	"log"
	"os"
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

	// Get the port from the environment variable, default to 8080
	port := os.Getenv("PORT") // Get the port from the environment variable
	if port == "" {           // If the environment variable is not set
		port = "8080" // Default to 8080
	}

	// Start server
	log.Printf("Starting server on :%s", port)
	if err := router.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
