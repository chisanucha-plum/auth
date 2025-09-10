package routes

import (
	handler "auth/internal/handlers"
	"auth/internal/middleware"
	"auth/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)

	// Initialize auth service and handler
	authService := services.NewAuthService(db)
	authHandler := handler.NewAuthHandler(authService)

	api := r.Group("/api/v1")

	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	protectedRoutes := api.Group("/protected")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.GET("/profile", authHandler.GetProfile)
	}

	return r
}
