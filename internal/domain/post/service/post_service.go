package service

import (
	"errors"
	"study-go-controller/internal/domain/post/entity"
	"study-go-controller/internal/domain/post/repository"
)

// PostService defines the contract for post business logic
type PostService interface {
	CreatePost(title, content string, authorID uint) (*entity.Post, error)
	GetPostByID(id uint) (*entity.Post, error)
	GetPostsByAuthorID(authorID uint) ([]*entity.Post, error)
	UpdatePost(id uint, title, content string, authorID uint) (*entity.Post, error)
	DeletePost(id uint, authorID uint) error
	GetAllPosts() ([]*entity.Post, error)
}

// postService implements PostService interface
type postService struct {
	postRepo repository.PostRepository
}

// NewPostService creates a new instance of PostService
func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{
		postRepo: postRepo,
	}
}

// CreatePost creates a new post
func (s *postService) CreatePost(title, content string, authorID uint) (*entity.Post, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	post := &entity.Post{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}

	if err := s.postRepo.Create(post); err != nil {
		return nil, err
	}

	// Fetch the post with author information
	return s.postRepo.GetByID(post.ID)
}

// GetPostByID retrieves a post by ID
func (s *postService) GetPostByID(id uint) (*entity.Post, error) {
	return s.postRepo.GetByID(id)
}

// GetPostsByAuthorID retrieves all posts by a specific author
func (s *postService) GetPostsByAuthorID(authorID uint) ([]*entity.Post, error) {
	return s.postRepo.GetByAuthorID(authorID)
}

// UpdatePost updates an existing post
func (s *postService) UpdatePost(id uint, title, content string, authorID uint) (*entity.Post, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if the user is the author of the post
	if post.AuthorID != authorID {
		return nil, errors.New("unauthorized: only the author can update this post")
	}

	if title == "" {
		return nil, errors.New("title is required")
	}

	post.Title = title
	post.Content = content

	if err := s.postRepo.Update(post); err != nil {
		return nil, err
	}

	return s.postRepo.GetByID(id)
}

// DeletePost deletes a post by ID
func (s *postService) DeletePost(id uint, authorID uint) error {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Check if the user is the author of the post
	if post.AuthorID != authorID {
		return errors.New("unauthorized: only the author can delete this post")
	}

	return s.postRepo.Delete(id)
}

// GetAllPosts retrieves all posts
func (s *postService) GetAllPosts() ([]*entity.Post, error) {
	return s.postRepo.GetAll()
}
