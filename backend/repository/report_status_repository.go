package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/dto"
	"geo_report_api/model"
	"log"

	"gorm.io/gorm"
)

func GetReportStatus(reportStatusID uint) (model.ReportStatus, error) {
	var reportStatus model.ReportStatus
	if err := config.Db.First(&reportStatus, reportStatusID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ReportStatus{}, nil
		}
		return model.ReportStatus{}, err
	}
	return reportStatus, nil
}

func GetReportStatusByName(status string) (uint, error) {
	var reportStatus model.ReportStatus
	if err := config.Db.Where("status = ?", status).First(&reportStatus).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return reportStatus.ID, nil
}

func CreateReportStatus(status string) (model.ReportStatus, error) {
	reportStatus := model.ReportStatus{Status: status}

	if err := config.Db.Where("status = ?", status).First(&reportStatus).Error; err == nil {
		return model.ReportStatus{}, errors.New("report type already exists")
	}

	if err := config.Db.Create(&reportStatus).Error; err != nil {
		log.Printf("Failed to create new report type: %v", err)
		return model.ReportStatus{}, errors.New("failed to create report type")
	}

	return reportStatus, nil
}

func GetAllreportStatus() ([]model.ReportStatus, error) {
	var reportStatuss []model.ReportStatus

	if err := config.Db.Find(&reportStatuss).Error; err != nil {
		return nil, err
	}

	return reportStatuss, nil
}

func GetStatusCounts() (dto.ChartData, error) {
	var chartData dto.ChartData
	var labels []string
	var data []int
	var backgroundColors []string

	rows, err := config.Db.Raw("SELECT report_statuses.status, COUNT(reports.id) AS report_count FROM report_statuses LEFT JOIN reports ON report_statuses.id = reports.report_status_id GROUP BY report_statuses.id").Rows()
	if err != nil {
		return chartData, err
	}
	defer rows.Close()

	for rows.Next() {
		var statusName string
		var reportCount int
		if err := rows.Scan(&statusName, &reportCount); err != nil {
			return chartData, err
		}
		labels = append(labels, statusName)
		data = append(data, reportCount)
		backgroundColors = append(backgroundColors, getRandomColor()) // Ensure getRandomColor is implemented
	}

	chartData = dto.ChartData{
		Labels: labels,
		Datasets: []dto.ReportTypeData{
			{
				Label:           "Total",
				Data:            data,
				BackgroundColor: backgroundColors,
				HoverOffset:     4,
			},
		},
	}

	return chartData, nil
}
