package service

import (
	"fmt"
	"geo_report_api/dto"
	"geo_report_api/repository"
	"geo_report_api/utils"
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

func ForgotPassword(email, newPassword string) error {
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to find user: %v", err)
	}

	err = repository.UpdateUserPassword(user, newPassword)
	if err != nil {
		return fmt.Errorf("failed to update user's password: %v", err)
	}

	err = utils.SendEmail(email, newPassword)
	if err != nil {
		return fmt.Errorf("failed to send password reset email: %v", err)
	}

	return nil
}
