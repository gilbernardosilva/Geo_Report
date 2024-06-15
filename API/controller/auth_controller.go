package controller

import (
	"geo_report_api/dto"
	"geo_report_api/service"

	"github.com/gin-gonic/gin"
)

// Login godoc
//
//	@Summary		Login
//	@Description	get token
//	@Tags			login
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginDTO	true	"query params"
//	@Success		200		{number}	token
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/user/login [post]
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
	c.JSON(200, gin.H{
		"message": "login user",
		"token":   t,
	})
}
