package controller

import (
	"geo_report_api/dto"
	"geo_report_api/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userLogin dto.LoginDTO
	err := c.ShouldBind(&userLogin)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	t, err := service.Login(userLogin)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "login user",
			"token":   t,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
	}
}
