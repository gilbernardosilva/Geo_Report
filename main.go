package main

import (
	"geo_report_api/config"
	"geo_report_api/controller"
	"geo_report_api/entities"
	"geo_report_api/utils"

	"github.com/gin-gonic/gin"
)

var Users []entities.User
var Reports []entities.Report

func main() {
	config.ConnectDB()

	defer config.CloseDB()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		adminRoutes := v1.Group("/admin")
		adminRoutes.Use(utils.JWTAuth())
		user := v1.Group("/user")
		{
			user.POST("/createUser", controller.CreateUser)
			user.GET("/:id", controller.GetUser)
		}

		report := v1.Group("/report")
		{
			report.GET("/:id", controller.GetReport)
		}

	}
	router.Run(":3000")
}
