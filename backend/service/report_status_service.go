package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"log"
)

func InsertReportStatus(status string) (model.ReportStatus, error) {
	if status == "" {
		return model.ReportStatus{}, errors.New("name can't be empty")
	}

	reportStatus, err := repository.CreateReportStatus(status)
	if err != nil {
		log.Printf("Failed to create report type: %v", err)
		return model.ReportStatus{}, errors.New("failed to create report type")
	}

	return reportStatus, nil
}

func GetAllReportStatus() ([]model.ReportStatus, error) {
	reportStatus, err := repository.GetAllreportStatus()
	if err != nil {
		return nil, err
	}

	return reportStatus, nil
}

func GetStatusCounts() (dto.ChartData, error) {
	chartData, err := repository.GetStatusCounts()
	if err != nil {
		return dto.ChartData{}, err
	}

	return chartData, nil
}
