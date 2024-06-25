package repository

import (
	"errors"
	"fmt"
	"geo_report_api/config"
	"geo_report_api/model"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var userNotFound = "user doesn't exist"

func CreateUser(user model.User) error {
	return config.Db.Create(&user).Error
}

func CheckIfUserExists(username string, email string) error {
	var user model.User

	if err := config.Db.Where("user_name = ?", username).First(&user).Error; err == nil {
		return fmt.Errorf("username already exists")
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking username: %v", err)
	}

	if err := config.Db.Where("email = ?", email).First(&user).Error; err == nil {
		return fmt.Errorf("email already exists")
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking email: %v", err)
	}

	return nil
}

func GetAllUsers() []model.User {
	var users []model.User
	config.Db.Find(&users)

	return users
}

func GetUser(userID uint64) (model.User, error) {
	var user model.User
	config.Db.First(&user, userID)
	if user.ID != 0 {
		return user, nil
	}
	return user, errors.New(userNotFound)
}

func UpdateUser(user model.User) (model.User, error) {
	if _, err := GetUser(user.ID); err == nil {
		config.Db.Save(&user)
		config.Db.Find(&user)
		return user, nil
	}
	return user, errors.New(userNotFound)
}

func DeleteUser(userID uint64) (bool, error) {
	if user, err := GetUser(userID); err == nil {
		config.Db.Delete(&user)
		return true, err
	}
	return false, errors.New(userNotFound)
}
