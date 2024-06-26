package controller

import (
	"fmt"
	"geo_report_api/dto"
	"geo_report_api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Report    Create Report
//
//	@Summary	Create Report
//	@Schemes
//	@Description	creates report in database
//	@Tags			Report
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.ReportCreatedDTO	true	"query params"
//	@Success		200		object	dto.ReportResponseDTO
//	@Failure		400		object	dto.ReportResponseDTO
//	@Router			/report [post]
func InsertReport(c *gin.Context) {
	var reportDTO dto.ReportCreatedDTO

	if err := c.ShouldBindJSON(&reportDTO); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Println("im here inserting")
	report, err := service.InsertReport(reportDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"report":  nil,
		})
		return
	}
	fmt.Println("im here inserted")

	err = service.AddPhotosToReport(reportDTO, report.ID, false)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"report":  nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Report created successfully",
		"report":  report,
	})
}

// Report    Update Report
//
//	@Summary	Update Report
//	@Schemes
//	@Description	updates report in database
//	@Tags			Report
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.ReportUpdateDTO	true	"Report Update Params"
//	@Success		200		object	dto.ReportResponseDTO
//	@Failure		400		object	dto.ReportResponseDTO
//	@Router			/report/ [put]
func EditReport(c *gin.Context) {
	var reportDTO dto.ReportUpdateDTO

	if err := c.ShouldBindJSON(&reportDTO); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"report":  nil,
		})
		return
	}

	report, err := service.UpdateReport(reportDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"report":  nil,
		})
	}

	c.JSON(200, gin.H{
		"message": "report successfully updated",
		"report":  report,
	})

}

// Report    Get Report
//
//	@Summary	Get Report
//	@Schemes
//	@Description	gets report from database
//	@Tags			Report
//	@Accept			json
//	@Produce		json
//	@Param			id	path	uint64	true	"report ID"
//	@Success		200	object	dto.ReportResponseDTO
//	@Failure		400	object	dto.ReportResponseDTO
//	@Router			/report/{id} [get]
func GetReport(c *gin.Context) {
	reportID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	report, err := service.GetReport(reportID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Report not found",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Report found",
		"report":  report,
	})
}

// GetReportsByUserID retrieves all reports for a specific user by user ID.
//
//	@Summary	Get Reports by User ID
//	@Schemes
//	@Description	Get all reports from a user by user ID
//	@Tags			Report
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64	true	"User ID"
//	@Success		200	{array}		model.Report
//	@Failure		404	{object}	dto.ReportResponseDTO
//	@Router			/report/user/{id} [get]
func GetReportsByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error parsing user id",
			"report":  nil,
		})
		return
	}

	reports, err := service.GetReportsByUserID(userID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"report":  nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Reports successfully retrieved",
		"reports": reports,
	})
}
