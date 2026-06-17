package model

import "time"

// Author represents a blog post author.
type Author struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Post represents a blog post.
type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Author      Author    `json:"author"`
	AuthorID    string    `json:"-"`
	ReadTime    string    `json:"readTime"`
	PublishedAt string    `json:"publishedAt"`
	Image       string    `json:"image"`
	ImageAlt    string    `json:"imageAlt"`
	Shares      int       `json:"shares"`
	Views       int       `json:"views"`
	IsFeatured  bool      `json:"isFeatured,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// CreatePostRequest is the DTO for creating a new post.
type CreatePostRequest struct {
	ID       string `json:"id"`
	Title    string `json:"title" validate:"required"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Category string `json:"category"`
	ReadTime string `json:"readTime"`
	Image    string `json:"image"`
	ImageAlt string `json:"imageAlt"`
	Status   string `json:"status"`
}

// UpdatePostRequest is the DTO for updating an existing post.
type UpdatePostRequest struct {
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	ReadTime    string `json:"readTime"`
	PublishedAt string `json:"publishedAt"`
	Image       string `json:"image"`
	ImageAlt    string `json:"imageAlt"`
	Shares      int    `json:"shares"`
	Views       int    `json:"views"`
	IsFeatured  *bool  `json:"isFeatured,omitempty"`
	Status      string `json:"status"`
}
