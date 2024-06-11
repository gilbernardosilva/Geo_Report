package entities

type Report struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	UserID      uint64 `gorm:"not null" json:"user_id"`
	User        User   `gorm:"foreign_key:UserID;constraint;onUpdate:CASCADE;onDelete:CASCADE" json:"user"`
	PhotoID     uint64 `gorm:"not null" json:"photo_id"`
	Photo       Photo  `gorm:"foreignkey:PhotoID;constraint;onUpdate:CASCADE;onDelete:CASCADE" json:"photo"`
	Description string `gorm:"type:text;not null" json:"description"`
}
