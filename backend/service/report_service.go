package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"log"

	"github.com/mashingan/smapping"
)

func InsertReport(reportDTO dto.ReportCreatedDTO) (model.Report, error) {
	var report model.Report

	err := smapping.FillStruct(&report, smapping.MapFields(&reportDTO))
	if err != nil {
		log.Printf("Failed to map report DTO to report struct: %v", err)
		return model.Report{}, errors.New("failed to map report")
	}

	exists, err := repository.ReportTypeExists(reportDTO.ReportTypeID)
	if !exists {
		log.Printf("Failed to find the Report type: %v", err)
		return model.Report{}, errors.New("Report Type doesn't exist")
	}

	if err := repository.InsertReport(report); err != nil {
		log.Printf("Failed to save report to database: %v", err)
		return model.Report{}, errors.New("failed to save report")
	}

	if len(reportDTO.Photos) > 0 {
		var photos []model.Photo
		for _, photoDTO := range reportDTO.Photos {
			var photo model.Photo
			err := smapping.FillStruct(&photo, smapping.MapFields(&photoDTO))
			if err != nil {
				return model.Report{}, errors.New("failed to map photo DTO to photo struct")
			}
			photos = append(photos, photo)
		}
		err = repository.AddPhotos(report.ID, photos)
		if err != nil {
			return model.Report{}, err
		}
	}

	return report, nil
}

func GetReport(id uint64) (model.Report, error) {
	if report, err := repository.GetReport(id); err == nil {
		return report, nil
	}
	return model.Report{}, errors.New("report does not exist")
}

func UpdateReport(reportID uint64, reportDTO dto.ReportUpdateDTO) (model.Report, error) {
	report, err := repository.GetReport(reportID)
	if err != nil {
		log.Printf("Failed to get report from db: %v", err)
		return model.Report{}, errors.New("report not found")
	}

	err = smapping.FillStruct(&report, smapping.MapFields(&reportDTO))
	if err != nil {
		log.Printf("Failed to map report DTO to report struct: %v", err)
		return model.Report{}, errors.New("fialed to map report")
	}

	exists, err := repository.ReportTypeExists(reportDTO.ReportTypeID)
	if err != nil {
		return model.Report{}, errors.New("failed to check issue type")
	}
	if !exists {
		return model.Report{}, errors.New("invalid issue type")
	}

	if err := repository.UpdateReport(report); err != nil {
		log.Printf("Failed to update report in db: %v", err)
		return model.Report{}, errors.New("failed to update report")
	}

	if len(reportDTO.Photos) > 0 {
		var photos []model.Photo
		for _, photoDTO := range reportDTO.Photos {
			var photo model.Photo
			err := smapping.FillStruct(&photo, smapping.MapFields(&photoDTO))
			if err != nil {
				return model.Report{}, errors.New("failed to map photo DTO to photo struct")
			}
			photos = append(photos, photo)
		}
		if err := repository.UpdatePhotos(reportID, photos); err != nil {
			log.Printf("Failed to update photos for report: %v", err)
			return model.Report{}, errors.New("failed to update photos for report")
		}
	}

	return report, nil
}
