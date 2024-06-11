package controller

import (
	"geo_report_api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
