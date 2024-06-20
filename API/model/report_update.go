package model

import (
	"time"
)

type ReportUpdate struct {
	ID             uint64       `gorm:"primary_key:auto_increment" json:"id"`
	ReportID       uint64       `gorm:"not null" json:"report_id"`
	Report         Report       `gorm:"foreignKey:ReportID" json:"report"`
	ReportStatusID uint         `gorm:"not null" json:"report_status_id"`
	ReportStatus   ReportStatus `gorm:"foreignKey:ReportStatusID" json:"report_status"`
	UpdatedAt      time.Time    `gorm:"type:timestamptz;not null" json:"updated_at"`
}
