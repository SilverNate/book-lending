package middleware

import (
	"book-lending-api/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(jwtService IJWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtService.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.APIResponse{
				Success: false, Error: "invalid or expired token",
			})
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}
