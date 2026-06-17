package repository

import "github.com/kemitbelajar/kemitbelajarblog-backend/internal/model"

// PostRepository defines the interface for post data access.
// This enables swapping implementations (e.g., PostgreSQL, in-memory, mock).
type PostRepository interface {
	// FindAll returns all posts ordered by creation date (newest first).
	FindAll() ([]model.Post, error)

	// FindByID returns a single post by its slug ID.
	FindByID(id string) (*model.Post, error)

	// FindByStatus returns posts filtered by status ("published" or "draft").
	FindByStatus(status string) ([]model.Post, error)

	// Create inserts a new post into the database.
	Create(post *model.Post) error

	// Update modifies an existing post.
	Update(post *model.Post) error

	// Delete removes a post by its slug ID.
	Delete(id string) error

	// IncrementShares increases the share count of a post by 1.
	IncrementShares(id string) error

	// IncrementViews increases the view count of a post by 1.
	IncrementViews(id string) error
}
