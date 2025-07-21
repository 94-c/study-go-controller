package repository

import (
	"study-go-controller/internal/domain/post/entity"

	"gorm.io/gorm"
)

// PostRepository defines the contract for post data operations
type PostRepository interface {
	Create(post *entity.Post) error
	GetByID(id uint) (*entity.Post, error)
	GetByAuthorID(authorID uint) ([]*entity.Post, error)
	Update(post *entity.Post) error
	Delete(id uint) error
	GetAll() ([]*entity.Post, error)
}

// postRepository implements PostRepository interface
type postRepository struct {
	db *gorm.DB
}

// NewPostRepository creates a new instance of PostRepository
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

// Create creates a new post in the database
func (r *postRepository) Create(post *entity.Post) error {
	return r.db.Create(post).Error
}

// GetByID retrieves a post by ID with author information
func (r *postRepository) GetByID(id uint) (*entity.Post, error) {
	var post entity.Post
	err := r.db.Preload("Author").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// GetByAuthorID retrieves all posts by a specific author
func (r *postRepository) GetByAuthorID(authorID uint) ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.db.Preload("Author").Where("author_id = ?", authorID).Find(&posts).Error
	return posts, err
}

// Update updates an existing post
func (r *postRepository) Update(post *entity.Post) error {
	return r.db.Save(post).Error
}

// Delete soft deletes a post by ID
func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Post{}, id).Error
}

// GetAll retrieves all posts with author information
func (r *postRepository) GetAll() ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.db.Preload("Author").Find(&posts).Error
	return posts, err
}
