package main

import (
	"geo_report_api/config"
	"geo_report_api/controller"
	"geo_report_api/entities"

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
		user := v1.Group("/user")
		{
			user.POST("/", controller.CreateUser)
			user.GET("/", controller.GetUsers)
			user.GET("/:id", controller.GetUser)
			user.PUT("/:id", controller.UpdateUser)
			user.DELETE("/:id", controller.DeleteUser)
			user.POST("/login", controller.Login)
		}

		report := v1.Group("/report")
		{
			report.POST("/", controller.CreateReport)
			report.GET("/", controller.GetReports)
			report.GET("/:id", controller.GetReport)
			report.PUT("/:id", controller.UpdateReport)
			report.DELETE("/:id", controller.DeleteReport)
		}

	}
	router.Run(":3000")
}
