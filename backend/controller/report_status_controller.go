package controller

import (
	"geo_report_api/service"

	"github.com/gin-gonic/gin"
)

func CreateReportStatus(c *gin.Context) {
	var request struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	reportStatus, err := service.InsertReportStatus(request.Status)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":      "report type created successfully",
		"reporty_type": reportStatus,
	})
}

func GetAllReportStatus(c *gin.Context) {
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

func GetStatusCounts(c *gin.Context) {
	chartData, err := service.GetStatusCounts()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to fetch status counts",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "status data for chart successfully fetched",
		"data":    chartData,
	})
}
