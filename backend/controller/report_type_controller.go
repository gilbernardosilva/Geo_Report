package controller

import (
	"geo_report_api/service"

	"github.com/gin-gonic/gin"
)

func CreateReportType(c *gin.Context) {
	var request struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	reportType, err := service.InsertReportType(request.Name)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":      "report type created successfully",
		"reporty_type": reportType,
	})
}

func GetAllReportTypes(c *gin.Context) {
	reportTypes, err := service.GetAllReportTypes()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to fetch all report types",
		})
	}

	c.JSON(200, gin.H{
		"message":      "report types fetched successfully",
		"report_types": reportTypes,
	})
}
