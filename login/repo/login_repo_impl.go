package repo

import (
	"LionChallenge/login"
	"LionChallenge/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

// LoginRepoImpl struct
type LoginRepoImpl struct {
	db *gorm.DB
}

// CreateLoginRepoImpl constructor
func CreateLoginRepoImpl(db *gorm.DB) login.LoginRepoInterface {
	return &LoginRepoImpl{db}
}

// GetPassByEmail function
func (l *LoginRepoImpl) GetPassByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := l.db.Where("email = ?", email).Find(&user).Error; err != nil {
		fmt.Errorf("[CreateLoginRepoImpl] get an error with : %w", err)
		return nil, err
	}
	return &user, nil

}
