package main

import (
	"log"
	"os"
	"study-go-controller/internal/domain/post/handler"
	postRepo "study-go-controller/internal/domain/post/repository"
	"study-go-controller/internal/domain/post/routes"
	postService "study-go-controller/internal/domain/post/service"
	userHandler "study-go-controller/internal/domain/user/handler"
	userRepo "study-go-controller/internal/domain/user/repository"
	userRoutes "study-go-controller/internal/domain/user/routes"
	userService "study-go-controller/internal/domain/user/service"
	"study-go-controller/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := db.AutoMigrate(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize repositories
	userRepository := userRepo.NewUserRepository(db.DB)
	postRepository := postRepo.NewPostRepository(db.DB)

	// Initialize services
	userSvc := userService.NewUserService(userRepository)
	postSvc := postService.NewPostService(postRepository)

	// Initialize handlers
	userHdl := userHandler.NewUserHandler(userSvc)
	postHdl := handler.NewPostHandler(postSvc)

	// Initialize Gin router
	router := gin.Default()

	// API version grouping
	v1 := router.Group("/api/v1")

	// Setup domain routes
	userRoutes.SetupUserRoutes(v1, userHdl)
	routes.SetupPostRoutes(v1, postHdl)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
