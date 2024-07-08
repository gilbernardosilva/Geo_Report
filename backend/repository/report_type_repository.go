package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/dto"
	"geo_report_api/model"
	"log"

	"gorm.io/gorm"
)

func GetReportType(reportTypeID uint64) (model.ReportType, error) {
	var reportType model.ReportType
	if err := config.Db.First(&reportType, reportTypeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ReportType{}, nil
		}
		return model.ReportType{}, err
	}
	return reportType, nil
}

func CreateReportType(name string) (model.ReportType, error) {
	reportType := model.ReportType{Name: name}

	if err := config.Db.Where("name = ?", name).First(&reportType).Error; err == nil {
		return model.ReportType{}, errors.New("report type already exists")
	}

	if err := config.Db.Create(&reportType).Error; err != nil {
		log.Printf("Failed to create new report type: %v", err)
		return model.ReportType{}, errors.New("failed to create report type")
	}

	return reportType, nil
}

func GetAllReportTypes() ([]model.ReportType, error) {
	var reportTypes []model.ReportType

	if err := config.Db.Find(&reportTypes).Error; err != nil {
		return nil, err
	}

	return reportTypes, nil
}

func GetReportTypeCounts() (dto.ChartData, error) {
	var chartData dto.ChartData

	rows, err := config.Db.Raw("SELECT report_types.name, COUNT(reports.id) AS report_count FROM report_types LEFT JOIN reports ON report_types.id = reports.report_type_id GROUP BY report_types.id").Rows()
	if err != nil {
		return chartData, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var count int

		if err := rows.Scan(&name, &count); err != nil {
			return chartData, err
		}

		chartData.Labels = append(chartData.Labels, name)
		chartData.Data = append(chartData.Data, count)
	}

	return chartData, nil
}
