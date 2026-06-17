package repository

import "github.com/kemitbelajar/kemitbelajarblog-backend/internal/model"

// UserRepository defines the interface for user-related database operations.
type UserRepository interface {
	// FindByUsername retrieves a user by their username.
	FindByUsername(username string) (*model.User, error)

	// FindByID retrieves a user by their ID.
	FindByID(id string) (*model.User, error)

	// FindAll retrieves all users.
	FindAll() ([]model.User, error)

	// Create adds a new user to the database.
	Create(user *model.User) error

	// Update modifies an existing user.
	Update(user *model.User) error

	// Delete removes a user by their ID.
	Delete(id string) error
}
