package main

import (
	"fiber-gorm/config"
	"fiber-gorm/models"
	"fiber-gorm/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
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

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Fiber + GORM CRUD API",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Health check endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Fiber + GORM CRUD API",
			"status":  "running",
		})
	})

	// Setup user routes
	routes.SetupUserRoutes(app)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	// Start server
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Server is running on port %s", port)
	if err := app.Listen(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
