package model

import (
	"geo_report_api/utils"
	"time"
)

type User struct {
	ID           uint64    `gorm:"primary_key:auto_increment" json:"id"`
	RoleID       uint      `gorm:"not null" json:"role_id"`
	Role         Role      `gorm:"foreignKey:RoleID" json:"role"`
	FirstName    string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName     string    `gorm:"type:varchar(255);not null" json:"last_name"`
	UserName     string    `gorm:"uniqueIndex;type:varchar(255);not null" json:"username"`
	Password     string    `gorm:"not null" json:"password"`
	Email        string    `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	CreationDate time.Time `gorm:"autoCreateTime"`
	Reports      []Report  `gorm:"foreignKey:UserID" json:"issues"`
}

func Hash(password string) (string, error) {
	return utils.CreateFromPassword(password)
}

func VerifyPassword(hashedPassword, password string) bool {
	return utils.ComparePasswordAndHash(password, hashedPassword)
}
