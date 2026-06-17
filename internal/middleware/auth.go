package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/response"
)

// JWTSecret should ideally be loaded from env.
var JWTSecret = []byte("kemitbelajar-super-secret-key")

// Protected ensures the route requires a valid JWT token.
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Failed(c, fiber.StatusUnauthorized, "Missing Authorization header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return response.Failed(c, fiber.StatusUnauthorized, "Invalid Authorization header format")
		}

		tokenString := parts[1]

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return JWTSecret, nil
		})

		if err != nil || !token.Valid {
			return response.Failed(c, fiber.StatusUnauthorized, "Invalid or expired token")
		}

		// Set claims in context
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Locals("userID", claims["id"])
			c.Locals("username", claims["username"])
		}

		return c.Next()
	}
}
