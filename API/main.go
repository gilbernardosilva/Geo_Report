package main

import (
	"geo_report_api/config"
	"geo_report_api/controller"
	"geo_report_api/entities"
	"geo_report_api/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Users []entities.User
var Reports []entities.Report

func main() {

	config.ConnectDB()

	defer config.CloseDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Allow sending cookies or Authorization headers
		MaxAge:           12 * time.Hour,
	}))

	v1 := router.Group("/api/v1")
	{
		adminRoutes := v1.Group("/admin")
		adminRoutes.Use(service.JWTAuth())

		adminUser := adminRoutes.Group("/user")
		{
			adminUser.GET("/:id", controller.GetUser)
			adminUser.GET("/", controller.GetAllUsers)
			adminUser.DELETE("/delete/:id", controller.DeleteUser)
		}

		user := v1.Group("/user")
		{
			user.POST("/register", controller.Register)
			user.POST("/login", controller.Login)
		}

		report := v1.Group("/report")
		{
			report.GET("/:id", controller.GetReport)
		}

	}
	router.Run(":8080")
}
