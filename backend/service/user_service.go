package service

import (
	"errors"
	"fmt"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"geo_report_api/utils"
	"log"
	"net/mail"

	"github.com/mashingan/smapping"
)

var userMappingFail = "Failed to map user DTO to user struct"

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
			RoleID:    user.RoleID,
			CreatedAt: user.CreationDate,
		}
		userDTOs = append(userDTOs, userDTO)
	}
	return userDTOs
}

// regular user start
func InsertUser(userDTO dto.UserCreatedDTO) (model.User, error) {
	var user model.User

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Printf(userMappingFail+": %v", err)
		return model.User{}, errors.New(userMappingFail)
	}

	// creates password hash
	user.Password, err = utils.CreateFromPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return model.User{}, errors.New("error hashing password")
	}

	// checks if user with unique info is already in db
	if err := repository.CheckIfUserExists(user.UserName, user.Email); err != nil {
		return model.User{}, errors.New("username or email already exists")
	}

	// checks if email is valid
	if err := valid(user.Email); err != nil {
		log.Printf("Invalid email format: %v", err)
		return model.User{}, errors.New("invalid email format")
	}

	// setting role ID to 0 since only regular users will use this function
	user.RoleID = 0

	createdUser, err := repository.CreateUser(&user)
	if err != nil {
		log.Printf("Failed to save user to database: %v", err)
		return model.User{}, errors.New("failed to save user")
	}

	return createdUser, nil
}

// regular user end

// admin user start
func InsertUserAdmin(userDTO dto.UserCreateAdminDTO) (model.User, error) {
	var user model.User

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Printf(userMappingFail+": %v", err)
		return model.User{}, errors.New(userMappingFail)
	}

	// creates password hash
	user.Password, err = utils.CreateFromPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return model.User{}, errors.New("error hashing password")
	}

	// checks if user with unique info is already in db
	if err := repository.CheckIfUserExists(user.UserName, user.Email); err != nil {
		return model.User{}, errors.New("username or email already exists")
	}

	// checks if email is valid
	if err := valid(user.Email); err != nil {
		log.Printf("Invalid email format: %v", err)
		return model.User{}, errors.New("invalid email format")
	}

	fmt.Println(userDTO)
	// check if role exists
	if err := repository.CheckIfRoleExists(user.RoleID); err != nil {
		log.Printf("Invalid role: %v", err)
		return model.User{}, errors.New("invalid role")
	}
	fmt.Println("checked role")

	// create user in the database
	createdUser, err := repository.CreateUser(&user)
	if err != nil {
		log.Printf("Failed to save user to database: %v", err)
		return model.User{}, errors.New("failed to save user")
	}

	return createdUser, nil
}

// AdminEditUser allows admins to edit user details, including changing their role
func AdminEditUser(userDTO dto.UserUpdateAdminDTO) (model.User, error) {
	var user model.User

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Printf(userMappingFail+": %v", err)
		return model.User{}, errors.New(userMappingFail)
	}

	// check if role exists
	if err := repository.CheckIfRoleExists(user.RoleID); err != nil {
		log.Printf("Invalid role: %v", err)
		return model.User{}, errors.New("invalid role")
	}

	// update user in the database
	updatedUser, err := repository.UpdateUser(user)
	if err != nil {
		log.Printf("Failed to update user in database: %v", err)
		return model.User{}, errors.New("failed to update user")
	}

	return updatedUser, nil
}

// admin user end

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
