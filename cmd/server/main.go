package main

import (
	"log"
	"os"
	"study-go-controller/pkg/container"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize DI container with automatic route registration
	c, err := container.NewContainer()
	if err != nil {
		log.Fatal("Failed to initialize container:", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// 🚀 Register all routes automatically
	c.RegisterRoutes(router)

	// Health check endpoint
	router.GET("/health", func(ctx *gin.Context) {
		registeredRoutes := c.GetRegisteredRoutes()
		ctx.JSON(200, gin.H{
			"status":       "ok",
			"message":      "Server is running with automatic routing",
			"total_routes": len(registeredRoutes),
			"auto_routes":  registeredRoutes,
		})
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Print registered routes for debugging
	routes := c.GetRegisteredRoutes()
	log.Println("\n🚀 ===== AUTOMATIC ROUTE REGISTRATION SUMMARY =====")
	for _, route := range routes {
		log.Printf("   %s %s", route.Method, route.Path)
	}
	log.Printf("📡 Total: %d routes automatically registered\n", len(routes))

	log.Printf("🌟 Server starting on port %s with automatic routing enabled!", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
