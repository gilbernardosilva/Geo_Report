package main

import (
	"geo_report_api/config"
	"geo_report_api/controller"
	"geo_report_api/entities"
	"geo_report_api/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "geo_report_api/docs"
)

var Users []entities.User
var Reports []entities.Report

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {

	config.ConnectDB()

	defer config.CloseDB()

	router := gin.Default()

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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":3000")
}
