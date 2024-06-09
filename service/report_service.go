package service

import (
	"geo_report_api/entities"
	"geo_report_api/repository"
)

func InsertReport(report entities.Report) entities.Report {
	repository.InsertReport(report)
	return report
}

func GetReport(id int) entities.Report {
	return repository.GetReport(id)
}
