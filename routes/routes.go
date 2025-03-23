package routes

import (
	"auth-service/internal/handlers"
	"auth-service/internal/models"
	"auth-service/package/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handlers.AuthHandler) *gin.Engine {
	r := gin.Default()

	// Public routes
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/validate", authHandler.ValidateToken)
		}
	}

	// Protected routes example (requiring admin role)
	protected := api.Group("/admin")
	protected.Use(middleware.AuthMiddleware(models.RoleAdmin))
	{
		// Add admin-only endpoints here
	}

	return r
}
