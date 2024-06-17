package controller

import (
	"geo_report_api/dto"
	"geo_report_api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "select users",
		"users":   service.GetAllUsers(),
	})
}

func Register(c *gin.Context) {
	var user dto.UserCreatedDTO
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	u, err := service.InsertUser(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"user":    "",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "user",
			"user":    u,
		})
	}

}

func GetUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := service.GetUser(userID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error (GetUser) - no user with that ID",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "select user",
		"user":    user,
	})

}

func DeleteUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	deleted, err := service.DeleteUser(userID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error (DeleteUser) - no user with that ID",
			"error":   err.Error(),
		})
	}
	if deleted {
		c.JSON(200, gin.H{
			"message": "user deleted successfully",
		})
	}
}
