package controller

import (
	"fmt"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/service"
	"strconv"
	"time"

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
	fmt.Println(reportDTO)
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

func GetAllReports(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid page parameter"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid limit parameter"})
		return
	}
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)

	areaIDStr := c.Query("area_id")
	var area model.Area
	if areaIDStr != "" {
		areaID, err := strconv.ParseUint(areaIDStr, 10, 64)
		if err != nil {
			c.JSON(500, gin.H{"error": "Invalid area ID"})
			return
		}

		area, err = service.GetAreaByID(areaID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Area not found"})
			return
		}
	}

	statusIDStr := c.Query("report_status_id")
	var statusID uint64 = 0

	if statusIDStr != "" {
		statusID, err = strconv.ParseUint(statusIDStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid status ID"})
			return
		}
	}

	reports, totalCount, err := service.GetAllReports(page, limit, startDate, endDate, area, statusID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"reports":     reports,
		"total_pages": int((totalCount + int64(limit) - 1) / int64(limit)),
	})
}

func UpdateReportStatus(c *gin.Context) {
	reportID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid report ID"})
		return
	}

	err = service.UpdateReportStatus(reportID, 5)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Report status updated successfully"})
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

	statusIDStr := c.Query("report_status_id")
	var statusID uint64 = 5

	if statusIDStr != "" {
		statusID, err = strconv.ParseUint(statusIDStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid status ID"})
			return
		}
	}

	reports, err := service.GetReportsByUserID(userID, statusID)
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
