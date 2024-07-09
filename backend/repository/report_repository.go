package repository

import (
	"errors"
	"fmt"
	"geo_report_api/config"
	"geo_report_api/model"
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

func GetAllReports(page, limit int, startDate, endDate time.Time, area *model.Area, statusID uint64) ([]model.Report, int64, error) {
	var reports []model.Report
	var totalCount int64

	query := config.Db.
		Preload("User").
		Preload("ReportStatus").
		Preload("ReportType").
		Preload("Photos", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos.id")
		})

	if !startDate.IsZero() {
		query = query.Where("report_date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("report_date <= ?", endDate)
	}
	if statusID != 5 {
		query = query.Where("report_status_id = ?", statusID)
	} else {
		query = query.Where("report_status_id != 5")
	}

	if err := query.Find(&reports).Error; err != nil {
		return nil, 0, fmt.Errorf("error retrieving reports: %v", err)
	}
	filteredReports := make([]model.Report, 0)
	if area != nil {
		centerPoint := fmt.Sprintf("ST_MakePoint(%f, %f)", area.Longitude, area.Latitude)
		query = query.Where("ST_DWithin(ST_MakePoint(longitude, latitude), ST_GeomFromText(?, 4326), ?)", centerPoint, area.Radius)
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
		var existingPhoto model.Photo
		if err := config.Db.Where("report_id = ? AND image = ?", reportID, photo.Image).First(&existingPhoto).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		} else {
			continue
		}

		// Create a new photo
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

func GetReportsByUserID(userID uint64, statusID uint64) ([]model.Report, error) {
	var reports []model.Report

	fmt.Println("getting report")
	query := config.Db.
		Preload("User.Role").
		Preload("Photos").
		Preload("ReportType").
		Preload("ReportStatus")

	if statusID != 5 {
		query = query.Where("user_id = ? AND report_status_id = ?", userID, statusID)
	} else {
		query = query.Where("user_id = ? AND report_status_id != 5", userID)
	}

	if err := query.Find(&reports).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Report{}, nil
		}
		return nil, err
	}

	return reports, nil
}
