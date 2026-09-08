package middleware

import (
	"net/http"
	"strings"

	"taskmanager/internal/utils"

	"github.com/gin-gonic/gin"
)

const userIDKey = "userID"

// AuthRequired validates the JWT in the Authorization header and stores the
// authenticated user's ID in the request context for downstream handlers.
func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			utils.Error(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			utils.Error(c, http.StatusUnauthorized, "Authorization header must be in the format: Bearer <token>")
			c.Abort()
			return
		}

		claims, err := utils.ParseJWT(parts[1], jwtSecret)
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set(userIDKey, claims.UserID)
		c.Next()
	}
}

// GetUserID reads the authenticated user's ID from the context.
// Must only be called on routes protected by AuthRequired.
func GetUserID(c *gin.Context) uint {
	id, _ := c.Get(userIDKey)
	userID, _ := id.(uint)
	return userID
}
