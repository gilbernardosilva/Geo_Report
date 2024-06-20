package model

import "time"

type Report struct {
	ID             uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Name           string         `gorm:"type:varchar(255);not null" json:"name"`
	UserID         uint64         `gorm:"not null" json:"user_id"`
	User           User           `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Photos         []Photo        `gorm:"foreignKey:ReportID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"photos"`
	Description    string         `gorm:"type:text;not null" json:"description"`
	IssueType      string         `gorm:"type:varchar(50);not null" json:"issue_type"`
	Latitude       float64        `gorm:"type:decimal(10,8);not null" json:"latitude"`
	Longitude      float64        `gorm:"type:decimal(11,8);not null" json:"longitude"`
	ReportStatusID uint           `gorm:"not null" json:"report_status_id"`
	ReportStatus   ReportStatus   `gorm:"foreignKey:ReportStatusID" json:"report_status"`
	ReportDate     time.Time      `gorm:"autoCreateTime" json:"report_date"`
	LastUpdate     time.Time      `gorm:"autoUpdateTime" json:"last_update"`
	Comments       []Comment      `gorm:"foreignKey:ReportID" json:"comments"`
	Updates        []ReportUpdate `gorm:"foreignKey:ReportID" json:"updates"`
}
