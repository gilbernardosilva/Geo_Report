package model

import (
	"log"

	"gorm.io/gorm"
)

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null;unique" json:"name"`
}

func InitRoles(db *gorm.DB) {
	roles := []Role{
		{ID: 0, Name: "user"},
		{ID: 1, Name: "authority"},
		{ID: 2, Name: "admin"},
	}

	// Loop through roles to insert them if they don't already exist
	for _, role := range roles {
		result := db.FirstOrCreate(&role, Role{ID: role.ID})
		if result.Error != nil {
			log.Fatalf("Failed to initialize role %s: %v", role.Name, result.Error)
		}
	}

	log.Println("Roles initialized successfully.")
}
