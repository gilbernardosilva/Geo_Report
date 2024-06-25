package config

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"geo_report_api/model"
)

var Db *gorm.DB

func ConnectDB() {

	err := godotenv.Load()
	fmt.Println(os.Getenv("DB_USERNAME"))
	// if os.Getenv("DB_USERNAME") == "" {
	// 	// Load from .env file (local development)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	fmt.Println("DB_HOST:", dbHost)
    fmt.Println("DB_PORT:", dbPort)
    fmt.Println("DB_USER:", dbUser)


	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbDatabase + " port=" + dbPort + " sslmode=disable"

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!" + dbDatabase)
	}

	err = Db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Report{}, &model.Area{}, &model.Authority{}, &model.Comment{},
		&model.Log{}, &model.Point{}, &model.ReportUpdate{}, &model.Role{}, &model.ReportStatus{})
	if err != nil {
		panic("Failed to migrate database!")
	}

}

func CloseDB() {
	db, err := Db.DB()
	if err != nil {
		panic("Failed to close database!")
	}
	db.Close()
}
