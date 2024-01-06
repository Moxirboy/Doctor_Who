package middleware

import (
	jwt "DoctorWho/internal/pkg/jwt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")

		if len(accessToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is not provided"})
			return
		}

		if !strings.HasPrefix(accessToken, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		// Extract the token from the header
		token := strings.TrimPrefix(accessToken, "Bearer ")

		claims, err := jwt.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			return
		}

		_, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid subject in token"})
			return
		}

		// You can add more validations or actions here if needed.

		// Set the claims in the context for later use
		c.Set("claims", claims)

		c.Next()
	}
}