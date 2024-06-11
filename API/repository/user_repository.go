package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/entities"
)

func InsertUser(user entities.User) entities.User {
	config.Db.Save(&user)
	config.Db.Find(&user)

	return user
}

func GetAllUsers() []entities.User {
	var users []entities.User
	config.Db.Find(&users)

	return users
}

func GetUser(userID uint64) (entities.User, error) {
	var user entities.User
	config.Db.First(&user, userID)
	if user.ID != 0 {
		return user, nil
	}
	return user, errors.New("user doesn't exist")
}

func UpdateUser(user entities.User) (entities.User, error) {
	if _, err := GetUser(user.ID); err == nil {
		config.Db.Save(&user)
		config.Db.Find(&user)
		return user, nil
	}
	return user, errors.New("user doesn't exist")
}

func UpdateUserToken(userID uint64, token string) (entities.User, error) {
	if user, err := GetUser(userID); err == nil {
		user.Token = token
		config.Db.Save(user)
		return user, nil
	}
	return entities.User{}, errors.New("user doesn't exist")
}

func DeleteUser(userID uint64) (bool, error) {
	if user, err := GetUser(userID); err == nil {
		config.Db.Delete(&user)
		return true, err
	}
	return false, errors.New("user doesn't exist")
}
