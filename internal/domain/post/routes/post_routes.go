package routes

import (
	"study-go-controller/internal/domain/post/handler"

	"github.com/gin-gonic/gin"
)

// SetupPostRoutes sets up routes for post domain
func SetupPostRoutes(router *gin.RouterGroup, postHandler *handler.PostHandler) {
	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("", postHandler.CreatePost)
		postRoutes.GET("", postHandler.GetAllPosts)
		postRoutes.GET("/:id", postHandler.GetPost)
		postRoutes.PUT("/:id", postHandler.UpdatePost)
		postRoutes.DELETE("/:id", postHandler.DeletePost)
		postRoutes.GET("/author/:id", postHandler.GetPostsByAuthor)
	}
}
