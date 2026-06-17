package response

import "github.com/gofiber/fiber/v2"

// APIResponse is the standardized JSON response envelope.
type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Success sends a 200 OK response with the standard format.
func Success(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(APIResponse{
		Status:  "success",
		Data:    data,
		Message: message,
	})
}

// Created sends a 201 Created response with the standard format.
func Created(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusCreated).JSON(APIResponse{
		Status:  "success",
		Data:    data,
		Message: message,
	})
}

// Failed sends an error response with the given HTTP status code.
func Failed(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Status:  "failed",
		Data:    nil,
		Message: message,
	})
}
