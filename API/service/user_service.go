package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"geo_report_api/utils"
	"log"
	"net/mail"

	"github.com/mashingan/smapping"
)

func GetAllUsers() []model.User {
	return repository.GetAllUsers()
}

func InsertUser(userDTO dto.UserCreatedDTO) (user model.User, error error) {
	user = model.User{}
	userResponse := dto.UserResponseDTO{}

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return
	}

	user.Password, err = utils.CreateFromPassword(user.Password)
	if err != nil {
		log.Fatal("error during pwd hash ", err)
		return
	}

	errDuplicated := repository.CheckIfUserExists(user.UserName, user.Email)

	if errDuplicated != nil {
		return
	}

	if !valid(user.Email) {
		return
	}

	user = repository.InsertUser(user)

	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		log.Fatal("failed to map response ", err)
		return model.User{}, err
	}

	return user, nil
}

func GetUser(userID uint64) (model.User, error) {
	if user, err := repository.GetUser(userID); err == nil {
		return user, nil
	}
	return model.User{}, errors.New("user doesn't exist")
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
