package handler

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/middleware"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/model"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/repository"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/response"
)

// PostHandler handles HTTP requests for blog posts.
type PostHandler struct {
	repo repository.PostRepository
}

// NewPostHandler creates a new PostHandler with the given repository.
func NewPostHandler(repo repository.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

// RegisterRoutes registers all post-related routes on the router.
func (h *PostHandler) RegisterRoutes(api fiber.Router) {
	api.Get("/posts", h.GetAllPosts)
	api.Get("/posts/:id", h.GetPost)
	
	// Protected routes
	api.Post("/posts", middleware.Protected(), h.CreatePost)
	api.Put("/posts/:id", middleware.Protected(), h.UpdatePost)
	api.Delete("/posts/:id", middleware.Protected(), h.DeletePost)
	
	// Engagement routes (public)
	api.Post("/posts/:id/share", h.SharePost)
	api.Post("/posts/:id/view", h.ViewPost)
}

// GetAllPosts returns all posts, optionally filtered by status query param.
// GET /api/posts?status=published
func (h *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	status := c.Query("status")

	var posts []model.Post
	var err error

	if status != "" {
		posts, err = h.repo.FindByStatus(status)
	} else {
		posts, err = h.repo.FindAll()
	}

	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to fetch posts")
	}

	return response.Success(c, posts, "Posts fetched successfully")
}

// GetPost returns a single post by ID.
// GET /api/posts/:id
func (h *PostHandler) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")

	post, err := h.repo.FindByID(id)
	if err != nil {
		log.Printf("Error fetching post %s: %v", id, err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to fetch post")
	}

	if post == nil {
		return response.Failed(c, fiber.StatusNotFound, "Post not found")
	}

	return response.Success(c, post, "Post fetched successfully")
}

// CreatePost creates a new blog post.
// POST /api/posts
func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var req model.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	// For a real app, use the authenticated user ID from c.Locals("userID")
	authorID := c.Locals("userID").(string)
	if authorID == "" {
		// Fallback for testing if middleware is disabled somehow, though it shouldn't be
		authorID = "cfb5435e-c288-4632-bae7-b76bb0619a6b"
	}

	if strings.TrimSpace(req.Title) == "" {
		return response.Failed(c, fiber.StatusBadRequest, "Title is required")
	}

	// Set defaults
	if req.Status == "" {
		req.Status = "draft"
	}
	if req.ID == "" {
		// Generate slug from title
		req.ID = generateSlug(req.Title)
	}

	post := &model.Post{
		ID:          req.ID,
		Title:       req.Title,
		Summary:     req.Summary,
		Content:     req.Content,
		Category:    req.Category,
		ReadTime:    req.ReadTime,
		PublishedAt: "Today",
		Image:       req.Image,
		ImageAlt:    req.ImageAlt,
		Shares:      0,
		Views:       0,
		IsFeatured:  false,
		Status:      req.Status,
	}

	if err := h.repo.Create(post); err != nil {
		log.Printf("Error creating post: %v", err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to create post")
	}

	return response.Created(c, post, "Post created successfully")
}

// UpdatePost updates an existing blog post.
// PUT /api/posts/:id
func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")

	// Get existing post
	existing, err := h.repo.FindByID(id)
	if err != nil {
		log.Printf("Error fetching post %s for update: %v", id, err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to fetch post")
	}
	if existing == nil {
		return response.Failed(c, fiber.StatusNotFound, "Post not found")
	}

	var req model.UpdatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Apply updates (only non-zero values)
	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.Summary != "" {
		existing.Summary = req.Summary
	}
	if req.Content != "" {
		existing.Content = req.Content
	}
	if req.Category != "" {
		existing.Category = req.Category
	}
	if req.ReadTime != "" {
		existing.ReadTime = req.ReadTime
	}
	if req.PublishedAt != "" {
		existing.PublishedAt = req.PublishedAt
	}
	if req.Image != "" {
		existing.Image = req.Image
	}
	if req.ImageAlt != "" {
		existing.ImageAlt = req.ImageAlt
	}
	if req.Shares != 0 {
		existing.Shares = req.Shares
	}
	if req.Views != 0 {
		existing.Views = req.Views
	}
	if req.IsFeatured != nil {
		existing.IsFeatured = *req.IsFeatured
	}
	if req.Status != "" {
		existing.Status = req.Status
	}

	if err := h.repo.Update(existing); err != nil {
		log.Printf("Error updating post %s: %v", id, err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to update post")
	}

	// Reload the full post to return with author data
	updated, err := h.repo.FindByID(id)
	if err != nil {
		return response.Success(c, existing, "Post updated successfully")
	}

	return response.Success(c, updated, "Post updated successfully")
}

// DeletePost removes a post by ID.
// DELETE /api/posts/:id
func (h *PostHandler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.repo.Delete(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return response.Failed(c, fiber.StatusNotFound, "Post not found")
		}
		log.Printf("Error deleting post %s: %v", id, err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to delete post")
	}

	return response.Success(c, nil, "Post deleted successfully")
}

// SharePost increments the share count of a post.
// POST /api/posts/:id/share
func (h *PostHandler) SharePost(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.repo.IncrementShares(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return response.Failed(c, fiber.StatusNotFound, "Post not found")
		}
		log.Printf("Error sharing post %s: %v", id, err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to share post")
	}

	// Return updated post
	post, err := h.repo.FindByID(id)
	if err != nil {
		return response.Success(c, nil, "Post shared successfully")
	}

	return response.Success(c, post, "Post shared successfully")
}

// ViewPost increments the view count of a post.
// POST /api/posts/:id/view
func (h *PostHandler) ViewPost(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.repo.IncrementViews(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return response.Failed(c, fiber.StatusNotFound, "Post not found")
		}
		log.Printf("Error viewing post %s: %v", id, err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to view post")
	}

	// Return updated post
	post, err := h.repo.FindByID(id)
	if err != nil {
		return response.Success(c, nil, "Post viewed successfully")
	}

	return response.Success(c, post, "Post viewed successfully")
}

// generateSlug creates a URL-friendly slug from a title.
func generateSlug(title string) string {
	slug := strings.ToLower(title)
	// Replace non-alphanumeric characters with hyphens
	var result strings.Builder
	for _, r := range slug {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			result.WriteRune(r)
		} else if r == ' ' || r == '-' || r == '_' {
			result.WriteRune('-')
		}
	}
	// Trim leading/trailing hyphens and collapse multiple hyphens
	s := result.String()
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	s = strings.Trim(s, "-")
	if s == "" {
		s = "untitled-post"
	}
	return s
}
