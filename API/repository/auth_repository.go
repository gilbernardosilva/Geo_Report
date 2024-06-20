package repository

import (
	"fmt"
	"geo_report_api/config"
	"geo_report_api/dto"
	"geo_report_api/model"
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
	return model.User{}, nil
}
