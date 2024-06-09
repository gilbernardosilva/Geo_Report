package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/entities"
)

func InsertUser(user entities.User) entities.User {
	config.Db.Save(&user)

	return user
}

func GetUser(userID uint64) (entities.User, error) {
	var user entities.User
	config.Db.Find(&user, userID)
	if user.ID != 0 {
		return user, nil
	}

	return user, errors.New("user does not exist")
}
