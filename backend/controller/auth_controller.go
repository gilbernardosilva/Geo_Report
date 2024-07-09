package controller

import (
	"crypto/rand"
	"fmt"
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
		c.JSON(400, gin.H{
			"message": err.Error(),
			"token":   "",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "login successful!",
			"token":   t,
		})
	}
}

// ResetPassword handles the HTTP POST request to reset a user's password
//
// @Summary Reset Password
// @Description Reset password for a user and send an email with a new password
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email address to reset password"
// @Success 200 {string} string "Password reset email sent successfully"
// @Failure 400 {object} string "Error message describing the issue"
// @Router /user/reset-password [post]
func ForgotPassword(c *gin.Context) {
	var resetPasswordDTO dto.ResetPasswordDTO
	if err := c.ShouldBindJSON(&resetPasswordDTO); err != nil {
		c.JSON(400, fmt.Sprintf("Invalid request format: %s", err.Error()))
		return
	}

	// Generate a new random password
	newPassword, err := generateRandomPassword(12)
	if err != nil {
		c.JSON(500, fmt.Sprintf("Failed to generate new password: %s", err.Error()))
		return
	}

	// Reset password logic
	err = service.ForgotPassword(resetPasswordDTO.Email, newPassword)
	if err != nil {
		c.JSON(500, fmt.Sprintf("Failed to reset password: %s", err.Error()))
		return
	}

	c.JSON(200, "Password reset email sent successfully")
}

func generateRandomPassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]"
	password := make([]byte, length)
	_, err := rand.Read(password)
	if err != nil {
		return "", fmt.Errorf("error generating random password: %v", err)
	}
	for i := range password {
		password[i] = charset[int(password[i])%len(charset)]
	}
	return string(password), nil
}
