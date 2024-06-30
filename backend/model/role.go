package model

import "gorm.io/gorm"

type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:false" json:"id"` // autoIncrement:false allows explicit ID setting
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func InitRoles(db *gorm.DB) {
	roles := []Role{
		{ID: 0, Name: "User"},
		{ID: 1, Name: "Authority"},
		{ID: 2, Name: "Admin"},
	}

	for _, role := range roles {
		var existingRole Role
		if result := db.FirstOrCreate(&existingRole, Role{ID: role.ID, Name: role.Name}); result.Error != nil {
			db.Create(&role)
		}
	}
}
