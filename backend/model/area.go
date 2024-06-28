package model

type Area struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name      string  `gorm:"type:varchar(255);not null" json:"name"`
	Latitude  float64 `gorm:"type:double precision;not null" json:"latitude"`
	Longitude float64 `gorm:"type:double precision;not null" json:"longitude"`
	Radius    float64 `gorm:"type:double precision;not null" json:"radius"`
}
