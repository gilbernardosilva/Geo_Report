package repository

import (
	"fmt"
	"geo_report_api/config"
	"geo_report_api/dto"
	"geo_report_api/entities"
)

func Login(loginDTO dto.LoginDTO) (entities.User, error) {
	var user entities.User
	err := config.Db.Where("user_name = ?", loginDTO.UserName).First(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	if !entities.VerifyPassword(user.Password, loginDTO.Password) {
		return entities.User{}, fmt.Errorf("invalid username or password")
	}
	return entities.User{}, nil
}
