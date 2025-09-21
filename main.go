package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"weathering-with-go/config"
	"weathering-with-go/handlers"
	"weathering-with-go/middleware"
	"weathering-with-go/services"
)

func main() {
	// Load configuration
	cfg := config.Load()
	
	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Set gin mode based on environment
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create weather service
	weatherService := services.NewWeatherService(cfg.OpenWeatherMapAPIKey)

	// Create gin router
	router := gin.Default()

	// Add middleware
	setupMiddleware(router, cfg)

	// Setup routes
	handlers.SetupRoutes(router, weatherService)

	// Start server
	log.Printf("Starting server on %s", cfg.GetServerAddress())
	log.Printf("Environment: %s", cfg.Environment)
	log.Printf("Log Level: %s", cfg.LogLevel)
	
	// Graceful shutdown setup
	go func() {
		if err := router.Run(cfg.GetServerAddress()); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	log.Println("Shutting down server...")
	log.Println("Server stopped")
}

// setupMiddleware configures middleware for the gin router
func setupMiddleware(router *gin.Engine, cfg *config.Config) {
	// Security headers
	router.Use(middleware.Security())
	
	// CORS middleware
	router.Use(middleware.CORS())

	// Request ID middleware
	router.Use(middleware.RequestID())

	// Logger middleware
	router.Use(middleware.Logger(cfg))

	// Recovery middleware
	router.Use(gin.Recovery())
}