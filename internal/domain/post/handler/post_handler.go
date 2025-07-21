package handler

import (
	"net/http"
	"strconv"
	"study-go-controller/internal/domain/post/dto"
	"study-go-controller/internal/domain/post/service"
	"study-go-controller/pkg/response"

	"github.com/gin-gonic/gin"
)

// PostHandler handles HTTP requests for post operations
// 🚀 자동 라우팅: 메서드 이름 기반으로 자동 등록됩니다!
type PostHandler struct {
	postService service.PostService
}

// NewPostHandler creates a new instance of PostHandler
func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

// CreatePost handles POST /posts
// 🔗 Auto Route: POST /api/v1/posts
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationErrorResponse(c, err)
		return
	}

	post, err := h.postService.CreatePost(req.Title, req.Content, req.AuthorID)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	postResponse := dto.ToPostResponse(post)
	response.SuccessResponse(c, http.StatusCreated, "Post created successfully", postResponse)
}

// GetPost handles GET /posts/:id
// 🔗 Auto Route: GET /api/v1/posts/:id
func (h *PostHandler) GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	postResponse := dto.ToPostResponse(post)
	response.SuccessResponse(c, http.StatusOK, "Post retrieved successfully", postResponse)
}

// GetAllPosts handles GET /posts
// 🔗 Auto Route: GET /api/v1/posts
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postResponses := dto.ToPostListResponseList(posts)
	response.SuccessResponse(c, http.StatusOK, "Posts retrieved successfully", postResponses)
}

// UpdatePost handles PUT /posts/:id
// 🔗 Auto Route: PUT /api/v1/posts/:id
func (h *PostHandler) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationErrorResponse(c, err)
		return
	}

	// TODO: Get author ID from authenticated user context
	authorID := uint(1) // Placeholder - should come from auth middleware

	post, err := h.postService.UpdatePost(uint(id), req.Title, req.Content, authorID)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	postResponse := dto.ToPostResponse(post)
	response.SuccessResponse(c, http.StatusOK, "Post updated successfully", postResponse)
}

// DeletePost handles DELETE /posts/:id
// 🔗 Auto Route: DELETE /api/v1/posts/:id
func (h *PostHandler) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	// TODO: Get author ID from authenticated user context
	authorID := uint(1) // Placeholder - should come from auth middleware

	if err := h.postService.DeletePost(uint(id), authorID); err != nil {
		response.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Post deleted successfully", nil)
}

// 🆕 GetPostsByAuthor handles GET /posts with author filter
// Note: This will be auto-registered as GET /posts/:id due to naming convention
// For custom routes, we might need a different approach
func (h *PostHandler) GetPostsByAuthor(c *gin.Context) {
	authorIDParam := c.Query("author_id")
	if authorIDParam == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "Author ID is required")
		return
	}

	authorID, err := strconv.ParseUint(authorIDParam, 10, 32)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid author ID")
		return
	}

	posts, err := h.postService.GetPostsByAuthorID(uint(authorID))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postResponses := dto.ToPostListResponseList(posts)
	response.SuccessResponse(c, http.StatusOK, "Posts retrieved successfully", postResponses)
}
