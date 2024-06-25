package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/model"
	"log"

	"gorm.io/gorm"
)

func ReportTypeExists(reportTypeID uint64) (bool, error) {
	var reportType model.ReportType
	if err := config.Db.First(&reportType, reportTypeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateReportType(name string) (model.ReportType, error) {
	reportType := model.ReportType{Name: name}

	if err := config.Db.Where("name- = ?", name).First(&reportType).Error; err == nil {
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
