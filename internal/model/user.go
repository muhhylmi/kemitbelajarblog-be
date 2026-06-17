package model

import "time"

// User represents an author/user in the system.
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Avatar       string    `json:"avatar"`
	Role         string    `json:"role"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // Never serialize the password hash
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// LoginRequest defines the payload for logging in.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse defines the payload returned upon successful authentication.
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CreateUserRequest defines the payload for creating a new user.
type CreateUserRequest struct {
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}
