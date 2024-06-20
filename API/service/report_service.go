package service

import (
	"errors"
	"geo_report_api/model"
	"geo_report_api/repository"
)

func InsertReport(report model.Report) model.Report {
	repository.InsertReport(report)
	return report
}

func GetReport(id uint64) (model.Report, error) {
	if book, err := repository.GetReport(id); err == nil {
		return book, nil
	}
	return model.Report{}, errors.New("report does not exist")
}
