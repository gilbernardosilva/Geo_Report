package repository

import (
	"geo_report_api/config"
	"geo_report_api/dto"
	"geo_report_api/entities"
	"geo_report_api/utils"
)

func Login(loginDTO dto.LoginDTO) (entities.User, error) {
	var user entities.User
	err := config.Db.Where("user_name = ?", loginDTO.UserName).First(&user).Error
	if err != nil {
		return entities.User{}, nil
	}
	//checking if password is the same as hashed one
	if utils.ComparePasswordAndHash(loginDTO.Password, user.Password) {
		return user, nil
	}
	return entities.User{}, nil
}
