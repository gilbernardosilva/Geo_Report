package model

type Point struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	AreaID    uint64 `gorm:"not null" json:"area_id"`
	Area      Area   `gorm:"foreignKey:AreaID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"area"`
	Latitude  string `gorm:"type:varchar(255);not null" json:"latitude"`
	Longitude string `gorm:"type:varchar(255);not null" json:"longitude"`
}
