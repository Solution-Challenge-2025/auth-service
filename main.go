package main

import (
	"auth-service/config"
	"auth-service/internal/handlers"
	"auth-service/internal/repositories"
	"auth-service/internal/services"
	"auth-service/migrations"
	"auth-service/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	log.Println("Running database migrations...")
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")

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
