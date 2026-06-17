package handler

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/middleware"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/model"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/repository"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/response"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

// RegisterAuthRoutes registers all auth routes.
func (h *AuthHandler) RegisterAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", h.Login)

	users := api.Group("/users")
	users.Get("/", h.ListUsers) // In a real app, this should be protected, but keeping it simple
	users.Post("/", middleware.Protected(), h.CreateUser)
	users.Put("/:id", middleware.Protected(), h.UpdateUser)
	users.Delete("/:id", middleware.Protected(), h.DeleteUser)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	user, err := h.userRepo.FindByUsername(req.Username)
	if err != nil {
		return response.Failed(c, fiber.StatusUnauthorized, "Invalid username or password")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return response.Failed(c, fiber.StatusUnauthorized, "Invalid username or password")
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(middleware.JWTSecret)
	if err != nil {
		return response.Failed(c, fiber.StatusInternalServerError, "Could not login")
	}

	return response.Success(c, model.AuthResponse{
		Token: t,
		User:  *user,
	}, "Login successful")
}

func (h *AuthHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userRepo.FindAll()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to fetch users")
	}

	return response.Success(c, users, "Users retrieved successfully")
}

func (h *AuthHandler) CreateUser(c *fiber.Ctx) error {
	// Check if current user is admin
	currentUserID := c.Locals("userID").(string)
	currentUser, err := h.userRepo.FindByID(currentUserID)
	if err != nil || (currentUser.Role != "Admin" && currentUser.Role != "Senior Editor") {
		return response.Failed(c, fiber.StatusForbidden, "Only Admins can create users")
	}

	var req model.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	if req.Username == "" || req.Password == "" || req.Name == "" {
		return response.Failed(c, fiber.StatusBadRequest, "Username, password, and name are required")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	user := &model.User{
		Name:         req.Name,
		Avatar:       req.Avatar,
		Role:         req.Role,
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	}

	// Default role/avatar if empty
	if user.Role == "" {
		user.Role = "Author"
	}
	if user.Avatar == "" {
		user.Avatar = "https://ui-avatars.com/api/?name=" + req.Username
	}

	if err := h.userRepo.Create(user); err != nil {
		log.Printf("Error creating user: %v", err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to create user (username might exist)")
	}

	return response.Created(c, user, "User created successfully")
}

func (h *AuthHandler) UpdateUser(c *fiber.Ctx) error {
	// Check if current user is admin
	currentUserID := c.Locals("userID").(string)
	currentUser, err := h.userRepo.FindByID(currentUserID)
	if err != nil || (currentUser.Role != "Admin" && currentUser.Role != "Senior Editor") {
		return response.Failed(c, fiber.StatusForbidden, "Only Admins can update users")
	}

	id := c.Params("id")
	userToUpdate, err := h.userRepo.FindByID(id)
	if err != nil {
		return response.Failed(c, fiber.StatusNotFound, "User not found")
	}

	var req model.CreateUserRequest // Reuse struct since it has the fields we need
	if err := c.BodyParser(&req); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	if req.Name != "" {
		userToUpdate.Name = req.Name
	}
	if req.Username != "" {
		userToUpdate.Username = req.Username
	}
	if req.Role != "" {
		userToUpdate.Role = req.Role
	}

	if err := h.userRepo.Update(userToUpdate); err != nil {
		log.Printf("Error updating user: %v", err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to update user")
	}

	return response.Success(c, userToUpdate, "User updated successfully")
}

func (h *AuthHandler) DeleteUser(c *fiber.Ctx) error {
	// Check if current user is admin
	currentUserID := c.Locals("userID").(string)
	currentUser, err := h.userRepo.FindByID(currentUserID)
	if err != nil || (currentUser.Role != "Admin" && currentUser.Role != "Senior Editor") {
		return response.Failed(c, fiber.StatusForbidden, "Only Admins can delete users")
	}

	id := c.Params("id")
	// Prevent deleting oneself
	if id == currentUserID {
		return response.Failed(c, fiber.StatusBadRequest, "Cannot delete your own account")
	}

	if err := h.userRepo.Delete(id); err != nil {
		if err.Error() == "user not found" {
			return response.Failed(c, fiber.StatusNotFound, "User not found")
		}
		log.Printf("Error deleting user: %v", err)
		return response.Failed(c, fiber.StatusInternalServerError, "Failed to delete user")
	}

	return response.Success(c, nil, "User deleted successfully")
}
