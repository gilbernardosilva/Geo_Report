package entities

import (
	"geo_report_api/utils"
)

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	RoleID    uint64 `gorm:"int" json:"roleID"`
	FirstName string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(255);not null" json:"last_name"`
	UserName  string `gorm:"uniqueIndex;type:varchar(255);not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
}

func Hash(password string) (string, error) {
	return utils.CreateFromPassword(password)
}

func VerifyPassword(hashedPassword, password string) bool {
	return utils.ComparePasswordAndHash(password, hashedPassword)
}
