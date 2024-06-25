package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/model"
)

func InsertReport(report model.Report) error {
	err := config.Db.Create(&report).Error
	config.Db.Preload("User").Find(&report)
	return err
}

func DeleteReport(report model.Report) {
	config.Db.Delete(&report)
}

func GetReport(reportID uint64) (model.Report, error) {
	var report model.Report
	config.Db.Preload("User").First(&report, reportID)
	config.Db.Preload("Photo").First(&report, reportID)
	if report.ID != 0 {
		return report, nil
	}

	return report, errors.New("report does not exist")
}

func UpdateReport(report model.Report) error {
	return config.Db.Save(&report).Error
}

func AddPhotos(reportID uint64, photos []model.Photo) error {
	for _, photo := range photos {
		photo.ReportID = reportID
		if err := config.Db.Create(&photo).Error; err != nil {
			return err
		}
	}
	return nil
}

func UpdatePhotos(reportID uint64, photos []model.Photo) error {
	if err := config.Db.Where("report_id = ?", reportID).Delete(&model.Photo{}).Error; err != nil {
		return err
	}

	for _, photo := range photos {
		photo.ReportID = reportID

		if err := config.Db.Create(&photo).Error; err != nil {
			return err
		}
	}

	return nil
}
