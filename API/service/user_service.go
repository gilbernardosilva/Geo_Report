package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/entities"
	"geo_report_api/repository"
	"geo_report_api/utils"
	"log"
	"net/mail"

	"github.com/mashingan/smapping"
)

func GetAllUsers() []entities.User {
	return repository.GetAllUsers()
}

func InsertUser(userDTO dto.UserCreatedDTO) (logintoken string, error error) {
	user := entities.User{}
	userResponse := dto.UserResponseDTO{}

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return "", err
	}

	user.Password, err = utils.CreateFromPassword(user.Password)
	if err != nil {
		log.Fatal("error during pwd hash ", err)
		return "", err
	}

	errDuplicated := repository.CheckIfUserExists(user.UserName, user.Email)

	if errDuplicated != nil {
		return "", errDuplicated
	}

	if !valid(user.Email) {
		return "", err
	}

	user = repository.InsertUser(user)

	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		log.Fatal("failed to map response ", err)
		return "", err
	}

	loginDTO := dto.LoginDTO{}

	loginDTO.UserName = user.UserName
	loginDTO.Password = user.Password

	logintoken, err = Login(loginDTO)

	if err != nil {
		log.Fatal("failed to log in using the new credentials", err)
		return "", err
	}

	return logintoken, nil
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

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
