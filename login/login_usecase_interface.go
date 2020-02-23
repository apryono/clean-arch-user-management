package login

import (
	"LionChallenge/model"
)

// LoginUsecaseInterface interface
type LoginUsecaseInterface interface {
	Login(user *model.User) (string, error)
	ValidatePassword(email string, pass string) (*model.User, error)
}
