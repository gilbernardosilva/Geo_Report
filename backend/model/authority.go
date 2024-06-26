package model

type Authority struct {
	ID             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID         uint64 `gorm:"not null" json:"user_id"`
	User           User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	AssignedAreaID uint64 `gorm:"not null" json:"assigned_area_id"`
	AssignedArea   Area   `gorm:"foreignKey:AssignedAreaID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"assigned_area"`
}
