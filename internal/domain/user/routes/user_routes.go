package routes

import (
	"study-go-controller/internal/domain/user/handler"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes sets up routes for user domain
func SetupUserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
