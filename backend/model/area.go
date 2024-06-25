package model

type Area struct {
	ID     uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name   string  `gorm:"type:varchar(255);not null" json:"name"`
	Points []Point `gorm:"foreignKey:AreaID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"points"`
}
