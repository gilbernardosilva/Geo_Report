package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"log"
)

func InsertReportType(name string) (model.ReportType, error) {
	if name == "" {
		return model.ReportType{}, errors.New("name can't be empty")
	}

	reportType, err := repository.CreateReportType(name)
	if err != nil {
		log.Printf("Failed to create report type: %v", err)
		return model.ReportType{}, errors.New("failed to create report type")
	}

	return reportType, nil
}

func GetAllReportTypes() ([]model.ReportType, error) {
	reportTypes, err := repository.GetAllReportTypes()
	if err != nil {
		return nil, err
	}

	return reportTypes, nil
}

func GetReportTypeCounts() (dto.ChartData, error) {
	chartData, err := repository.GetReportTypeCounts()
	if err != nil {
		return dto.ChartData{}, err
	}

	return chartData, nil
}
