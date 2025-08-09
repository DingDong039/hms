package middleware

import (
	"strings"

	"github.com/DingDong039/hms/internal/models"
	"github.com/DingDong039/hms/internal/services"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates a middleware for JWT authentication
func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, models.NewErrorResponse(401, "authorization header is required"))
			return
		}

		// Check if the header format is valid
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, models.NewErrorResponse(401, "invalid authorization header format"))
			return
		}

		// Extract the token
		tokenString := parts[1]

		// Validate the token
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, models.NewErrorResponse(401, "invalid or expired token"))
			return
		}

		// Set user information in the context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
