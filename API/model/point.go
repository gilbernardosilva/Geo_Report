package model

type Point struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	AreaID    uint64  `gorm:"not null" json:"area_id"`
	Area      Area    `gorm:"foreignKey:AreaID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"area"`
	Latitude  float64 `gorm:"type:decimal(10,8);not null" json:"latitude"`
	Longitude float64 `gorm:"type:decimal(11,8);not null" json:"longitude"`
}
