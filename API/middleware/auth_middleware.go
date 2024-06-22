package middleware

import (
	"fmt"
	"geo_report_api/config"
	"geo_report_api/model"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var allowedRoles = map[string][]uint64{
	"/api/v1/dashboard":        {0, 1, 2}, // Regular users can access
	"/api/v1/reports":          {0},       // Regular users can access
	"/api/v1/reports/create":   {0, 1, 2}, // Regular users can create
	"/api/v1/reports/:id":      {0, 1},    // Both users and authorities can view a specific report
	"/api/v1/reports/:id/edit": {2},       // Only admin can edit
	"/api/v1/authority":        {1, 2},    // Only authorities can access this endpoint
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid token claims"})
			return
		}

		userID, ok := claims["id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID in token"})
			return
		}

		var user model.User
		if err := config.Db.First(&user, uint(userID)).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		path := c.Request.URL.Path

		if !isAuthorized(uint64(user.RoleID), path) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}

		c.Set("user", user)
		c.Next()
	}

}

func isAuthorized(roleID uint64, path string) bool {
	allowed, exists := allowedRoles[path]
	if !exists {
		return false
	}
	for _, r := range allowed {
		if r == roleID {
			return true
		}
	}
	return false
}
