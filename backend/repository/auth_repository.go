package repository

import (
	"fmt"
	"geo_report_api/config"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/utils"
)

func Login(loginDTO dto.LoginDTO) (model.User, error) {
	var user model.User
	err := config.Db.Where("user_name = ?", loginDTO.UserName).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	if !model.VerifyPassword(user.Password, loginDTO.Password) {
		return model.User{}, fmt.Errorf("invalid username or password")
	}
	return user, nil
}

func FindUserByEmail(email string) (*model.User, error) {
	user := &model.User{}

	if err := config.Db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return user, nil
}

func UpdateUserPassword(user *model.User, newPassword string) error {
	user.Password, _ = utils.CreateFromPassword(newPassword)

	if err := config.Db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user's password: %v", err)
	}

	fmt.Printf("Updated password for user %s", user.Email)
	return nil
}
