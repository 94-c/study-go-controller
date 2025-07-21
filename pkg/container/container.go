package container

import (
	"log"
	"study-go-controller/internal/domain/post/handler"
	postRepo "study-go-controller/internal/domain/post/repository"
	postService "study-go-controller/internal/domain/post/service"
	userHandler "study-go-controller/internal/domain/user/handler"
	userRepo "study-go-controller/internal/domain/user/repository"
	userService "study-go-controller/internal/domain/user/service"
	"study-go-controller/pkg/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Container holds all dependencies
type Container struct {
	DB *gorm.DB

	// Repositories
	UserRepo userRepo.UserRepository
	PostRepo postRepo.PostRepository

	// Services
	UserService userService.UserService
	PostService postService.PostService

	// Handlers
	UserHandler *userHandler.UserHandler
	PostHandler *handler.PostHandler

	// Auto Router
	AutoRouter *AutoRouter
}

// NewContainer creates and initializes all dependencies
func NewContainer() (*Container, error) {
	// Initialize database
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	// Run migrations
	if err := db.AutoMigrate(); err != nil {
		return nil, err
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

	// Initialize auto router
	autoRouter := NewAutoRouter()

	container := &Container{
		DB:          db.DB,
		UserRepo:    userRepository,
		PostRepo:    postRepository,
		UserService: userSvc,
		PostService: postSvc,
		UserHandler: userHdl,
		PostHandler: postHdl,
		AutoRouter:  autoRouter,
	}

	// 🚀 자동으로 모든 핸들러 라우트 등록
	container.registerAllHandlers()

	return container, nil
}

// registerAllHandlers automatically registers all handler routes
func (c *Container) registerAllHandlers() {
	log.Println("🔄 Starting automatic route registration...")

	// Register User domain routes
	c.AutoRouter.RegisterHandler("/users", c.UserHandler)

	// Register Post domain routes
	c.AutoRouter.RegisterHandler("/posts", c.PostHandler)

	log.Println("✅ Automatic route registration completed!")
}

// RegisterRoutes registers all domain routes automatically
func (c *Container) RegisterRoutes(router *gin.Engine) {
	// API version grouping
	v1 := router.Group("/api/v1")

	// 🚀 자동으로 모든 라우트 등록
	c.AutoRouter.RegisterRoutes(v1)

	log.Printf("📡 Total registered routes: %d", len(c.AutoRouter.GetRoutes()))
}

// GetRegisteredRoutes returns all registered routes for debugging
func (c *Container) GetRegisteredRoutes() []RouteInfo {
	return c.AutoRouter.GetRoutes()
}
