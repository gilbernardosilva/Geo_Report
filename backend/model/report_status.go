package model

import (
	"log"

	"gorm.io/gorm"
)

type ReportStatus struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Status string `gorm:"type:varchar(100);not null" json:"status"`
}

func InitReportStatus(db *gorm.DB) {
	reportStatuses := []ReportStatus{
		{Status: "Pending"},
		{Status: "In Progress"},
		{Status: "Completed"},
		{Status: "Rejected"},
	}

	for _, status := range reportStatuses {
		var existingStatus ReportStatus
		if err := db.Where("status = ?", status.Status).First(&existingStatus).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&status).Error; err != nil {
					log.Fatalf("Failed to initialize report status %s: %v", status.Status, err)
				}
			} else {
				log.Fatalf("Failed to check report status %s: %v", status.Status, err)
			}
		}
	}

	log.Println("Report statuses initialized successfully.")
}
