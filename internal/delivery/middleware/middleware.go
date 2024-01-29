package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware that checks if a user is authenticated (has a user_id in the session).
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// Check if user_id exists in the session
		userID := session.Get("userId")
		
		if userID == nil {
			c.Status(401)
			// Redirect to the login page if user_id is not present
			c.Redirect(http.StatusFound, "/v1/login")
			c.Abort()
			return
		}
		id:=userID.(int)
		user:=strconv.Itoa(id)
		c.SetCookie("userId",  user, 3600, "/", "localhost", false, true)

	
		// Call the next handler in the chain
		c.Next()
	}
}
