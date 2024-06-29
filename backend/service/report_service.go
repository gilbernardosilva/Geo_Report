package service

import (
	"errors"
	"fmt"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"log"
	"time"

	"github.com/mashingan/smapping"
)

var reportError = "Report Status doesn't exist"

func InsertReport(reportDTO dto.ReportCreatedDTO) (model.Report, error) {
	report := model.Report{
		Name:           reportDTO.Name,
		UserID:         reportDTO.UserID,
		Description:    reportDTO.Description,
		ReportTypeID:   reportDTO.ReportTypeID,
		Latitude:       reportDTO.Latitude,
		Longitude:      reportDTO.Longitude,
		ReportStatusID: reportDTO.ReportStatusID,
	}

	if len(reportDTO.Photos) > 0 {
		report.Photos = make([]model.Photo, len(reportDTO.Photos))
		for i, photoDTO := range reportDTO.Photos {
			report.Photos[i] = model.Photo{
				Image: photoDTO,
			}
		}
	}

	// Check for report type existence
	if _, err := repository.GetReportType(report.ReportTypeID); err != nil {
		return model.Report{}, fmt.Errorf("report type doesn't exist: %v", err)
	}

	// Check for user existence
	if _, err := repository.GetUser(report.UserID); err != nil {
		return model.Report{}, fmt.Errorf("user doesn't exist: %v", err)
	}

	// Check for report status existence
	if _, err := repository.GetReportStatus(report.ReportStatusID); err != nil {
		return model.Report{}, fmt.Errorf("report status doesn't exist: %v", err)
	}

	if report, _ = repository.InsertReport(report); report.ID == 0 {
		log.Printf("Failed to save report to database")
		return model.Report{}, errors.New("failed to save report")
	}

	return report, nil
}

func AddPhotosToReport(reportDTO dto.ReportCreatedDTO, reportID uint64, update bool) error {
	if len(reportDTO.Photos) > 0 {
		var newPhotos []model.Photo
		for _, photoString := range reportDTO.Photos {
			newPhoto := model.Photo{
				ReportID: reportID,
				Image:    photoString,
			}
			newPhotos = append(newPhotos, newPhoto)
		}

		if update {
			if err := repository.UpdatePhotos(reportID, newPhotos); err != nil {
				log.Printf("Failed to update photos for report: %v", err)
				return errors.New("failed to update photos for report")
			}
		} else {
			err := repository.AddPhotos(reportID, newPhotos)
			if err != nil {
				return errors.New("failed to add photos to report")
			}
		}
	}
	return nil
}

func UpdateReport(reportDTO dto.ReportUpdateDTO) (model.Report, error) {
	report, err := repository.GetReport(reportDTO.ID)
	if err != nil {
		log.Printf("Failed to get report from db: %v", err)
		return model.Report{}, errors.New("report not found")
	}

	err = smapping.FillStruct(&report, smapping.MapFields(&reportDTO))
	if err != nil {
		log.Printf("Failed to map report DTO to report struct: %v", err)
		return model.Report{}, errors.New("fialed to map report")
	}

	_, err = repository.GetReportType(reportDTO.ReportTypeID)
	if err != nil {
		log.Printf("Failed to find the Report type: %v", err)
		return model.Report{}, errors.New("Report Type doesn't exist")
	}

	_, err = repository.GetUser(report.UserID)
	if err != nil {
		log.Printf("User doesn't exist: %v", err)
		return model.Report{}, errors.New("User ID doesn't exist")
	}

	_, err = repository.GetReportStatus(reportDTO.ReportStatusID)
	if err != nil {
		log.Printf("reportError")
		return model.Report{}, errors.New(reportError)
	}

	if err := repository.UpdateReport(report); err != nil {
		log.Printf("Failed to update report in db: %v", err)
		return model.Report{}, errors.New("failed to update report")
	}

	return report, nil
}

func GetReport(id uint64) (model.Report, error) {
	if report, err := repository.GetReport(id); err == nil {
		return report, nil
	}
	return model.Report{}, errors.New("report does not exist")
}

func GetAllReports(page, limit int, startDate, endDate time.Time, area model.Area) ([]model.Report, int64, error) {
	return repository.GetAllReports(page, limit, startDate, endDate, &area)
}

func GetReportsByUserID(userID uint64) ([]model.Report, error) {
	_, err := repository.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return repository.GetReportsByUserID(userID)
}

func UpdateReportStatus(reportID uint64, newStatusID uint) error {
	err := repository.UpdateReportStatus(reportID, newStatusID)
	if err != nil {
		return fmt.Errorf("error updating report status: %v", err)
	}
	return nil
}
