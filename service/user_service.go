package service

import (
	"errors"
	"geo_report_api/entities"
	"geo_report_api/repository"
)

func GetUser(id uint64) (entities.User, error) {
	if user, err := repository.GetUser(id); err == nil {
		return user, nil
	}
	return entities.User{}, errors.New("user does not exist")
}
