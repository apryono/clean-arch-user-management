package usecase

import (
	"LionChallenge/login"
	"LionChallenge/model"
	"LionChallenge/utils"
	"errors"
)

// LoginUsecaseImpl struct
type LoginUsecaseImpl struct {
	loginRepo login.LoginRepoInterface
}

// CreateLoginUsecaseImpl function
func CreateLoginUsecaseImpl(loginRepo login.LoginRepoInterface) login.LoginUsecaseInterface {
	return &LoginUsecaseImpl{loginRepo}
}

// Login function
func (l *LoginUsecaseImpl) Login(user *model.User) (string, error) {
	user, err := l.ValidatePassword(user.Email, user.Password)
	if err != nil {
		utils.LogError("Login", "Error when validate password", err)
		return "", err
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.LogError("Login", "Error when generate token", err)
		return "", err
	}
	return token, nil
}

// ValidatePassword function
func (l *LoginUsecaseImpl) ValidatePassword(email string, pass string) (*model.User, error) {
	user, err := l.loginRepo.GetPassByEmail(email)
	if err != nil {
		utils.LogError("ValidatePassword", "Error when get password", err)
		return nil, err
	}
	valid := utils.ComparePassword(user.Password, []byte(pass))
	if !valid {
		return nil, errors.New("Password is wrong")
	}
	return user, nil
}
