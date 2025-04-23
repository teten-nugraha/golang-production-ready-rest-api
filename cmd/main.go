package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/config"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/handlers"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/middlewares"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/repositories"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/services"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/utils"
	"github.com/teten-nugraha/golang-production-ready-rest-api/pkg/logging"
	"go.uber.org/zap"
)

// @title Golang REST API
// @version 1.0
// @description This is a sample REST API with Go.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Load environment variables
	config.LoadEnv()
	cfg := config.NewConfig()

	// Initialize logger
	logging.InitLogger()

	// Initialize Redis
	utils.InitRedis(cfg.RedisAddr)
	defer utils.RedisClient.Close()

	// Initialize database (example with GORM)
	db, err := utils.InitDB(cfg)
	if err != nil {
		logging.Logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JwtSecret)
	userService := services.NewUserService(userRepo)

	// Initialize controllers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// Create Gin router
	r := gin.Default()

	// Setup Swagger
	config.SetupSwagger(r)

	// Middleware
	r.Use(middlewares.CORSMiddleware(cfg.CORSAllowOrigin))
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.RateLimitMiddleware(utils.RedisClient, cfg.RateLimit))

	// API Versioning
	v1 := r.Group("/api/v1")

	// Auth routes
	v1.POST("/auth/register", authHandler.Register)
	v1.POST("/auth/login", authHandler.Login)

	// Protected routes
	protected := v1.Group("")
	protected.Use(middlewares.AuthMiddleware(cfg.JwtSecret))
	{
		protected.GET("/users", userHandler.GetAllUsers)
		protected.GET("/users/:id", userHandler.GetUserByID)
		protected.PUT("/users/:id", userHandler.UpdateUser)
		protected.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, utils.NewSuccessResponse(nil, "API is healthy"))
	})

	// Start server
	r.Run(":" + cfg.AppPort)
}
