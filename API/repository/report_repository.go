package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/model"
)

func InsertReport(report model.Report) (model.Report, error) {
	result := config.Db.Create(&report)
	if result.Error != nil {
		return model.Report{}, result.Error
	}

	err := config.Db.Preload("User.Role").Preload("Photos.Report").Preload("ReportType").Preload("ReportStatus").First(&report, report.ID).Error
	if err != nil {
		return model.Report{}, err
	}
	return report, nil
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

func GetReportsByUserID(userID uint64) ([]model.Report, error) {
	var reports []model.Report
	err := config.Db.Preload("User.Role").Preload("Photos").Preload("ReportType").Preload("ReportStatus").
		Where("user_id = ?", userID).Find(&reports).Error

	if err != nil {
		return nil, err
	}

	return reports, nil
}
