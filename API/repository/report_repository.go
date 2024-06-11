package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/entities"
)

func InsertReport(report entities.Report) entities.Report {
	config.Db.Save(&report)
	config.Db.Preload("User").Find(&report)
	return report
}

func DeleteReport(report entities.Report) {
	config.Db.Delete(&report)
}

func GetReport(reportID uint64) (entities.Report, error) {
	var report entities.Report
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
