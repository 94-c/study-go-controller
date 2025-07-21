package enums

// PostStatus represents the status of a post
type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
	PostStatusArchived  PostStatus = "archived"
	PostStatusDeleted   PostStatus = "deleted"
)

// String returns the string representation of PostStatus
func (s PostStatus) String() string {
	return string(s)
}

// IsValid checks if the post status is valid
func (s PostStatus) IsValid() bool {
	switch s {
	case PostStatusDraft, PostStatusPublished, PostStatusArchived, PostStatusDeleted:
		return true
	default:
		return false
	}
}

// GetAllPostStatuses returns all valid post statuses
func GetAllPostStatuses() []PostStatus {
	return []PostStatus{
		PostStatusDraft,
		PostStatusPublished,
		PostStatusArchived,
		PostStatusDeleted,
	}
}

// PostCategory represents categories for posts
type PostCategory string

const (
	PostCategoryTech      PostCategory = "tech"
	PostCategoryLifestyle PostCategory = "lifestyle"
	PostCategoryNews      PostCategory = "news"
	PostCategoryReview    PostCategory = "review"
)

// String returns the string representation of PostCategory
func (c PostCategory) String() string {
	return string(c)
}

// IsValid checks if the post category is valid
func (c PostCategory) IsValid() bool {
	switch c {
	case PostCategoryTech, PostCategoryLifestyle, PostCategoryNews, PostCategoryReview:
		return true
	default:
		return false
	}
}
