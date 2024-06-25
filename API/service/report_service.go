package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"log"

	"github.com/mashingan/smapping"
)

var reportError = "Report Status doesn't exist"

func InsertReport(reportDTO dto.ReportCreatedDTO) (model.Report, error) {
	var report model.Report

	err := smapping.FillStruct(&report, smapping.MapFields(&reportDTO))
	if err != nil {
		log.Printf("Failed to map report DTO to report struct: %v", err)
		return model.Report{}, errors.New("failed to map report")
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
		log.Printf(reportError)
		return model.Report{}, errors.New(reportError)
	}

	if report, err = repository.InsertReport(report); err != nil {
		log.Printf("Failed to save report to database: %v", err)
		return model.Report{}, errors.New("failed to save report")
	}

	return report, nil
}

func AddPhotosToReport(reportDTO dto.ReportCreatedDTO, reportID uint64, update bool) error {

	if len(reportDTO.Photos) > 0 {
		var photos []model.Photo
		for _, photoDTO := range reportDTO.Photos {
			var photo model.Photo
			err := smapping.FillStruct(&photo, smapping.MapFields(&photoDTO))
			if err != nil {
				return errors.New("failed to map photo DTO to photo struct")
			}
			photo.ReportID = reportID
			photos = append(photos, photo)
		}

		if update {
			if err := repository.UpdatePhotos(reportID, photos); err != nil {
				log.Printf("Failed to update photos for report: %v", err)
				return errors.New("failed to update photos for report")
			}
		} else {
			err := repository.AddPhotos(reportID, photos)
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

func GetReportsByUserID(userID uint64) ([]model.Report, error) {
	_, err := repository.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return repository.GetReportsByUserID(userID)
}
