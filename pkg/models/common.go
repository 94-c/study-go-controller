package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel contains common fields for all entities
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// PaginationRequest represents pagination parameters
type PaginationRequest struct {
	Page     int `json:"page" form:"page" binding:"min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}

// PaginationResponse represents pagination response
type PaginationResponse struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalCount int64       `json:"total_count"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

// SortRequest represents sorting parameters
type SortRequest struct {
	SortBy  string `json:"sort_by" form:"sort_by"`
	SortDir string `json:"sort_dir" form:"sort_dir" binding:"oneof=asc desc"`
}
