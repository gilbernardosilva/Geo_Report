package model

type ReportType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`
}
