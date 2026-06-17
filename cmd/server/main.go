package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/config"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/database"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/handler"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/repository"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	// Resolve migrations directory relative to the project root
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..", "..")
	migrationsDir := filepath.Join(projectRoot, "backend", "migrations")

	if err := database.RunMigrations(db, migrationsDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Repositories
	postRepo := repository.NewPostgresPostRepository(db)
	userRepo := repository.NewPostgresUserRepository(db)

	// Handlers
	postHandler := handler.NewPostHandler(postRepo)
	authHandler := handler.NewAuthHandler(userRepo)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Kemitbelajar Blog API",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"data":    fiber.Map{"service": "kemitbelajar-blog-api"},
			"message": "Service is healthy",
		})
	})

	// Register routes under /api
	api := app.Group("/api")
	postHandler.RegisterRoutes(api)
	authHandler.RegisterAuthRoutes(api)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Kemitbelajar Blog API starting on %s", addr)
	log.Printf("📡 CORS allowed origins: %s", cfg.AllowedOrigins)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
