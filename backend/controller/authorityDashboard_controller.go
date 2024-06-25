package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorityDashboard(c *gin.Context) {
	user, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user from context"})
		return
	}

	c.JSON(http.StatusOK, user)
}
