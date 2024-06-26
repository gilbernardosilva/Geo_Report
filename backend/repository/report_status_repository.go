package repository

import (
	"errors"
	"geo_report_api/config"
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
