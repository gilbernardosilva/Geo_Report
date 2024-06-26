package model

import "time"

type Comment struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	ReportID  uint64    `gorm:"not null" json:"report_id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
