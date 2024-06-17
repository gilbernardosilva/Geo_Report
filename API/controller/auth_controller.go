package controller

import (
	"geo_report_api/dto"
	"geo_report_api/service"

	"github.com/gin-gonic/gin"
)

// User	Login
//
//	@Summary	Login
//	@Schemes
//	@Description	returns token for login
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginDTO	true	"query params"
//	@Success		200		{object}	dto.LoginResponse
//	@Router			/user/login [post]
func Login(c *gin.Context) {
	var userLogin dto.LoginDTO
	err := c.ShouldBind(&userLogin)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"token":   "",
		})
		return
	}

	t, err := service.Login(userLogin)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
			"token":   "",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "login user",
			"token":   t,
		})
	}

}
