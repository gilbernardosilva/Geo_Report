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

func GetAllUsers() []dto.UserResponseDTO {
	users := repository.GetAllUsers()
	var userDTOs []dto.UserResponseDTO
	for _, user := range users {
		userDTO := dto.UserResponseDTO{
			ID:        user.ID,
			UserName:  user.UserName,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			RoleName:  user.Role.Name,
			CreatedAt: user.CreationDate,
		}
		userDTOs = append(userDTOs, userDTO)
	}
	return userDTOs
}

func InsertUser(userDTO dto.UserCreatedDTO) (new model.User, error error) {
	var user model.User

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Printf("Failed to map user DTO to user struct: %v", err)
		return model.User{}, errors.New("failed to map user")
	}

	user.Password, err = utils.CreateFromPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return model.User{}, errors.New("error hashing password")
	}

	if err := repository.CheckIfUserExists(user.UserName, user.Email); err != nil {
		return model.User{}, errors.New("username or email already exists")
	}

	if err := valid(user.Email); err != nil {
		log.Printf("Invalid email format: %v", err)
		return model.User{}, errors.New("invalid email format")
	}

	if err := repository.CreateUser(user); err != nil {
		log.Printf("Failed to save user to database: %v", err)
		return model.User{}, errors.New("failed to save user")
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

func EditUser(user model.User) (model.User, error) {
	if user, err := repository.UpdateUser(user); err == nil {
		return user, nil
	}
	return model.User{}, errors.New("error editing user")
}

func valid(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
