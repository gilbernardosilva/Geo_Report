package controller

import (
	"net/http"

	"geo_report_api/model"

	"github.com/gin-gonic/gin"
)

// Dashboard godoc
// @Summary User dashboard
// @Description Returns the user dashboard
// @Tags users
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Router /dashboard [get]
func Dashboard(c *gin.Context) {
	userValue, exists := c.Get("user") // Get user from context
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user from context"})
		return
	}

	user, ok := userValue.(model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type in context"})
		return
	}

	if user.RoleID != 0 { // Check the user's role
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, user)
}
