package dto

import (
	"study-go-controller/internal/domain/post/entity"
	userDto "study-go-controller/internal/domain/user/dto"
	"time"
)

// CreatePostRequest represents the request body for creating a post
type CreatePostRequest struct {
	Title    string `json:"title" binding:"required,min=1,max=200"`
	Content  string `json:"content" binding:"max=10000"`
	AuthorID uint   `json:"author_id" binding:"required"`
}

// UpdatePostRequest represents the request body for updating a post
type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"max=10000"`
}

// PostResponse represents the response body for post data
type PostResponse struct {
	ID        uint                  `json:"id"`
	Title     string                `json:"title"`
	Content   string                `json:"content"`
	AuthorID  uint                  `json:"author_id"`
	Author    *userDto.UserResponse `json:"author,omitempty"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

// PostListResponse represents a simplified post response for lists
type PostListResponse struct {
	ID        uint                  `json:"id"`
	Title     string                `json:"title"`
	AuthorID  uint                  `json:"author_id"`
	Author    *userDto.UserResponse `json:"author,omitempty"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

// ToPostResponse converts Post entity to PostResponse DTO
func ToPostResponse(post *entity.Post) *PostResponse {
	response := &PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	// Include author information if available
	if post.Author.ID != 0 {
		response.Author = userDto.ToUserResponse(&post.Author)
	}

	return response
}

// ToPostListResponse converts Post entity to PostListResponse DTO (without content)
func ToPostListResponse(post *entity.Post) *PostListResponse {
	response := &PostListResponse{
		ID:        post.ID,
		Title:     post.Title,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	// Include author information if available
	if post.Author.ID != 0 {
		response.Author = userDto.ToUserResponse(&post.Author)
	}

	return response
}

// ToPostResponseList converts slice of Post entities to slice of PostResponse DTOs
func ToPostResponseList(posts []*entity.Post) []*PostResponse {
	responses := make([]*PostResponse, len(posts))
	for i, post := range posts {
		responses[i] = ToPostResponse(post)
	}
	return responses
}

// ToPostListResponseList converts slice of Post entities to slice of PostListResponse DTOs
func ToPostListResponseList(posts []*entity.Post) []*PostListResponse {
	responses := make([]*PostListResponse, len(posts))
	for i, post := range posts {
		responses[i] = ToPostListResponse(post)
	}
	return responses
}
