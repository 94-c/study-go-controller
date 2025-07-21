package entity

import (
	userEntity "study-go-controller/internal/domain/user/entity"
	"time"

	"gorm.io/gorm"
)

// Post represents the post entity in the domain
type Post struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	Title     string          `json:"title" gorm:"not null"`
	Content   string          `json:"content" gorm:"type:text"`
	AuthorID  uint            `json:"author_id" gorm:"not null"`
	Author    userEntity.User `json:"author" gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}

// TableName returns the table name for Post entity
func (Post) TableName() string {
	return "posts"
}
