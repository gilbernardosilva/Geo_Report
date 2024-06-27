package main

import (
	"geo_report_api/config"
	"geo_report_api/controller"
	"geo_report_api/middleware"
	"geo_report_api/model"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "geo_report_api/docs"
)

var Users []model.User
var Reports []model.Report

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

	model.InitRoles(config.Db)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := router.Group("/api/v1")
	{
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/dashboard", controller.Dashboard)
			protected.GET("/authority", controller.AuthorityDashboard)

			protected.GET("/reportTypes", controller.GetAllReportTypes)
			protected.GET("/reportStatus", controller.GetAllReportStatus)

			userProtected := protected.Group("/user")
			{
				userProtected.PUT("/edit/:id", controller.EditUser)
			}
			// admin routes
			adminRoutes := protected.Group("/admin")
			{
				adminRoutes.GET("")
			}
			adminUser := adminRoutes.Group("/user")
			{
				adminUser.GET("/:id", controller.GetUser)
				adminUser.GET("", controller.GetAllUsers)
				adminUser.DELETE("/delete/:id", controller.DeleteUser)
			}

			adminUser = adminRoutes.Group("/report")
			{
				adminUser.POST("/reportType", controller.CreateReportType)
				adminUser.POST("/reportStatus", controller.CreateReportStatus)

			}

			//authority routes

			// report routes
			report := protected.Group("/report")
			{
				report.GET("/:id", controller.GetReport)
				report.GET("/user/:id", controller.GetReportsByUserID)
				report.POST("", controller.InsertReport)
				report.PUT("", controller.EditReport)
			}

		}

		public := v1.Group("/user")
		{
			public.POST("/register", controller.Register)
			public.POST("/login", controller.Login)
		}

		public = v1.Group("/area")
		{
			public.GET("/:id/point", controller.GetPointsByAreaID)
			public.POST("/:id/point", controller.CreatePoint)
			public.GET("", controller.GetAllAreas)
			public.POST("", controller.CreateArea)
			public.GET("/:id", controller.GetAreaByID)
			public.PUT("/:id", controller.UpdateArea)
			public.DELETE("/:id", controller.DeleteArea)
		}

		public = v1.Group("/point")
		{
			public.PUT("/:id", controller.UpdatePoint)
			public.DELETE("/:id", controller.DeletePoint)
		}

	}

	//dev - remove later
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

}