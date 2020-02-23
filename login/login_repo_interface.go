package login

import "LionChallenge/model"

// LoginRepoInterface interface
type LoginRepoInterface interface {
	GetPassByEmail(email string) (*model.User, error)
}
