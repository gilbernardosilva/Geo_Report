package repository

import (
	"errors"
	"fmt"
	"geo_report_api/config"
	"geo_report_api/model"
	"geo_report_api/utils"
	"math"
	"time"

	"gorm.io/gorm"
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
	result := config.Db.Preload("User").Where("report_status_id != ?", 5).First(&report, reportID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return report, errors.New("report does not exist")
		}
		return report, fmt.Errorf("error retrieving report: %v", result.Error)
	}

	if err := config.Db.Preload("Photos").First(&report, reportID).Error; err != nil {
		return report, fmt.Errorf("error retrieving photos for report %d: %v", reportID, err)
	}

	return report, nil
}

func GetAllReports(page, limit int, startDate, endDate time.Time, area *model.Area) ([]model.Report, int64, error) {
	var reports []model.Report
	var totalCount int64

	query := config.Db.
		Preload("User").
		Preload("ReportStatus").
		Preload("ReportType").
		Preload("Photos", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos.id")
		}).
		Where("report_status_id != ?", 5)

	if !startDate.IsZero() {
		query = query.Where("report_date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("report_date <= ?", endDate)
	}

	if err := query.Find(&reports).Error; err != nil {
		return nil, 0, fmt.Errorf("error retrieving reports: %v", err)
	}
	filteredReports := make([]model.Report, 0)
	for _, report := range reports {
		if area != nil {
			distance := utils.HaversineDistance(
				area.Latitude, area.Longitude, report.Latitude, report.Longitude,
			)
			if distance <= area.Radius {
				filteredReports = append(filteredReports, report)
			}
		} else {
			filteredReports = append(filteredReports, report)
		}
	}
	totalCount = int64(len(filteredReports))

	startIndex := (page - 1) * limit
	endIndex := int(math.Min(float64(startIndex+limit), float64(totalCount)))

	if area.ID != 0 {
		paginatedReports := filteredReports[startIndex:endIndex]
		return paginatedReports, totalCount, nil
	} else {
		return reports, totalCount, nil
	}
}

func UpdateReport(report model.Report) error {
	return config.Db.Save(&report).Error
}

func UpdateReportStatus(reportID uint64, newStatusID uint) error {
	result := config.Db.Model(&model.Report{}).Where("id = ?", reportID).Update("report_status_id", newStatusID)
	return result.Error
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
	err := config.Db.
		Preload("User.Role").
		Preload("Photos").
		Preload("ReportType").
		Preload("ReportStatus").
		Where("user_id = ? AND report_status_id != ?", userID, 5).
		Find(&reports).
		Error

	if err != nil {
		return nil, err
	}

	return reports, nil
}
