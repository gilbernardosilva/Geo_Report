package model

type ReportStatus struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Status string `gorm:"type:varchar(100);not null" json:"status"`
}
