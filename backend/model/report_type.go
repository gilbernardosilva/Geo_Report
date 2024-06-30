package model

import (
	"log"

	"gorm.io/gorm"
)

type ReportType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`
}

func InitReportType(db *gorm.DB) {
	reportTypes := []ReportType{
		{Name: "Pothole"},
		{Name: "Street Light Outage"},
		{Name: "Traffic Signal Issue"},
		{Name: "Road Debris"},
		{Name: "Flooding"},
		{Name: "Sign Damage"},
		{Name: "Overgrown Vegetation"},
	}

	for _, reportType := range reportTypes {
		var existingReportType ReportType
		if err := db.Where("name = ?", reportType.Name).First(&existingReportType).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&reportType).Error; err != nil {
					log.Fatalf("Failed to initialize report type %s: %v", reportType.Name, err)
				}
			} else {
				log.Fatalf("Failed to check report type %s: %v", reportType.Name, err)
			}
		}
	}

	log.Println("Report types initialized successfully.")
}
