package model

import "time"

type Photo struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	ReportID  uint64    `gorm:"not null" json:"report_id"`
	Report    Report    `gorm:"foreignKey:ReportID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"report"`
	URL       string    `gorm:"type:varchar(255);not null" json:"url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
