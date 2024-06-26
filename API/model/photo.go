package model

import "time"

type Photo struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	ReportID  uint64    `gorm:"not null"`
	Report    Report    `gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Image     string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
