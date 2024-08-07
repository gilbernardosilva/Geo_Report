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

	// initializing the db
	model.InitRoles(config.Db)
	model.InitReportStatus(config.Db)
	model.InitReportType(config.Db)

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

			// admin routes
			adminRoutes := protected.Group("/admin")
			{
				adminRoutes.GET("")
			}

			userProtected := protected.Group("user")
			{
				userProtected.PUT("/edit/:id", controller.EditUser)
			}

			adminUser := adminRoutes.Group("/user")
			{
				adminUser.GET("/:id", controller.GetUser)
				adminUser.GET("", controller.GetAllUsers)
				adminUser.POST("", controller.AdminRegister)
				adminUser.PUT("", controller.AdminEditUser)
				adminUser.DELETE("/delete/:id", controller.DeleteUser)
			}

			adminUser = adminRoutes.Group("/report")
			{
				adminUser.POST("/reportType", controller.CreateReportType)
				adminUser.POST("/reportStatus", controller.CreateReportStatus)
				adminUser.GET("", controller.GetAllReports)
				adminUser.PUT("/delete/:id", controller.DeleteReport)
				adminUser.PUT("/:id", controller.UpdateReportStatus)
			}
			// area routes
			area := adminRoutes.Group("/area")
			{
				area.GET("", controller.GetAllAreas)
				area.POST("", controller.CreateArea)
				area.GET("/:id", controller.GetAreaByID)
				area.PUT("/:id", controller.UpdateArea)
				area.DELETE("/:id", controller.DeleteArea)
			}
			point := protected.Group("/point")
			{
				point.PUT("/:id", controller.UpdatePoint)
				point.DELETE("/:id", controller.DeletePoint)
			}

			// report routes
			report := protected.Group("/report")
			{
				report.GET("/:id", controller.GetReport)
				report.GET("/user/:id", controller.GetReportsByUserID)
				report.POST("", controller.InsertReport)
				report.PUT("", controller.EditReport)
				report.GET("/status/chart", controller.GetStatusCounts)
				report.GET("/types/chart", controller.GetReportTypeCounts)
			}

		}

		public := v1.Group("/user")
		{
			public.POST("/register", controller.Register)
			public.POST("/login", controller.Login)
			public.POST("/forgotPassword", controller.ForgotPassword)
		}

	}

	//dev - remove later
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

}
