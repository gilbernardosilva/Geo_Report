package controller

import (
	"fmt"
	"geo_report_api/dto"
	"geo_report_api/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User	Register User
//
//	@Summary	Register
//	@Schemes
//	@Description	registers user in database
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserCreatedDTO	true "query params"
//	@Success		200		{object}	dto.UserResponseDTO
//	@Failure		400		{object}	dto.UserResponseDTO
//	@Router			/user/register [post]
func Register(c *gin.Context) {
	var user dto.UserCreatedDTO
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	userCreated, err := service.InsertUser(user)
	if err != nil {
		// Log the error for debugging purposes
		log.Printf("Error creating user: %v", err)
		c.JSON(400, gin.H{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}

	token, err := service.GenerateJWT(userCreated)
	if err != nil {
		// Log the error for debugging purposes
		log.Printf("Error generating JWT: %v", err)
		c.JSON(400, gin.H{
			"message": "Failed to generate token",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"token":   token,
	})

}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "select users",
		"users":   service.GetAllUsers(),
	})
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

// User    Edit User
//
//	@Summary	Edit User
//	@Schemes
//	@Description	edits user information in database
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserUpdateDTO	true "query params"
//	@Success		200		{object}	dto.UserResponseDTO
//	@Failure		400		{object}	dto.UserResponseDTO
//	@Router			/user/edit [put]
func EditUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := service.GetUser(userID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
			"token":   "",
		})
		return
	}

	userValue, exists := c.Get("userID")
	fmt.Println(userValue)
	if !exists {
		c.JSON(500, gin.H{"error": "Failed to retrieve user from context"})
		return
	}

	if userID != uint64(userValue.(float64)) {
		c.JSON(404, gin.H{"error": "You are not authorized to edit this profile"})
		return
	}
	var updatedUserData dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&updatedUserData); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user.FirstName = updatedUserData.FirstName
	user.LastName = updatedUserData.LastName
	user.UserName = updatedUserData.UserName
	user.Email = updatedUserData.Email

	user, err = service.EditUser(user)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
			"token":   "",
		})
		return
	}

	token, err := service.GenerateJWT(user)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
			"token":   "",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "EditUser was successful",
		"token":   token,
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
