package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/entities"
	"geo_report_api/repository"
	"log"

	"github.com/mashingan/smapping"
)

func GetAllUsers() []entities.User {
	return repository.GetAllUsers()
}

func InsertUser(userDTO dto.UserCreatedDTO) (dto.UserResponseDTO, error) {
	user := entities.User{}
	userResponse := dto.UserResponseDTO{}

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return userResponse, err
	}

	user = repository.InsertUser(user)

	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		log.Fatal("failed to map response ", err)
		return userResponse, err
	}

	return userResponse, nil
}

func GetUser(userID uint64) (entities.User, error) {
	if user, err := repository.GetUser(userID); err == nil {
		return user, nil
	}
	return entities.User{}, errors.New("user doesn't exist")
}

func DeleteUser(userID uint64) (bool, error) {
	if deleted, err := repository.DeleteUser(userID); err == nil {
		return deleted, nil
	}
	return false, errors.New("user doesn't exist")
}
