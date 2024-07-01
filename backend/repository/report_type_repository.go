package repository

import (
	"errors"
	"fmt"
	"geo_report_api/config"
	"geo_report_api/dto"
	"geo_report_api/model"
	"log"
	"math/rand"

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
	var labels []string
	var data []int
	var backgroundColors []string

	rows, err := config.Db.Raw("SELECT report_types.name, COUNT(reports.id) AS report_count FROM report_types LEFT JOIN reports ON report_types.id = reports.report_type_id GROUP BY report_types.id").Rows()
	if err != nil {
		return chartData, err
	}
	defer rows.Close()

	for rows.Next() {
		var typeName string
		var reportCount int
		if err := rows.Scan(&typeName, &reportCount); err != nil {
			return chartData, err
		}
		labels = append(labels, typeName)
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

func getRandomColor() string {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}
