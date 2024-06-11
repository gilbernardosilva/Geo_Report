package entities

import "time"

type Photo struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Caption   string    `gorm:"type:varchar(255)" json:"caption"`
	PhotoUrl  string    `gorm:"type:varchar(255)" json:"photo_url"`
	ReportID  int       `gorm: "foreignKey:ReportID" json:"report_id"`
	CreatedAt time.Time `gorm: "DateTime" json:"created_at"`
	UpdatedAt time.Time `gorm: "DateTime" json:"updated_at"`
}
