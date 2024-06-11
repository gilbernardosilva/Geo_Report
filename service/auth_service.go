package service

import (
	"geo_report_api/dto"
	"geo_report_api/repository"
)

func Login(loginDTO dto.LoginDTO) (string, error) {
	token := ""
	user, err := repository.Login(loginDTO)
	if err != nil {
		return token, err
	}

	token, _ = GenerateJWT(user)

	return token, nil
}
