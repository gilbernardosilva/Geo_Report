package controller

import (
	"geo_report_api/dto"
	"geo_report_api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertReport(c *gin.Context) {
	var reportDTO dto.ReportCreatedDTO

	if err := c.ShouldBindJSON(&reportDTO); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid JSON data",
		})
		return
	}

	report, err := service.InsertReport(reportDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Report created successfully",
		"report":  report,
	})
}

func GetReport(c *gin.Context) {
	reportID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	report, err := service.GetReport(reportID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Report found",
		"report":  report,
	})
}
