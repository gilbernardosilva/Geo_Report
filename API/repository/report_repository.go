package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/model"
)

func InsertReport(report model.Report) model.Report {
	config.Db.Save(&report)
	config.Db.Preload("User").Find(&report)
	return report
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

func UpdateReport() {

}

func SelectReport() {

}
