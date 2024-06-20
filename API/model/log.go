package model

import (
	"time"
)

type Log struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Action    string    `gorm:"type:varchar(255);not null" json:"action"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
