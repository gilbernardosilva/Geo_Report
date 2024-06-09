package service

import (
	"errors"
	"geo_report_api/entities"
	"geo_report_api/repository"
)

func InsertReport(report entities.Report) entities.Report {
	repository.InsertReport(report)
	return report
}

func GetReport(id uint64) (entities.Report, error) {
	if book, err := repository.GetReport(id); err == nil {
		return book, nil
	}
	return entities.Report{}, errors.New("report does not exist")
}
