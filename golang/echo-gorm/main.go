package main

import (
	"echo-gorm/config"
	"echo-gorm/models"
	"echo-gorm/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize database connection
	if err := config.InitDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Auto migrate the schema
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database synchronized")

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Health check endpoint
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "Echo + GORM CRUD API",
			"status":  "running",
		})
	})

	// Setup user routes
	routes.SetupUserRoutes(e)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	// Start server
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Server is running on port %s", port)
	if err := e.Start(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
