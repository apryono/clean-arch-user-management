package repo

import (
	"LionChallenge/model"
	"LionChallenge/user"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type UserRepoImpl struct {
	db *gorm.DB
}

// CreateUserRepoImpl constructor for connection to database
func CreateUserRepoImpl(db *gorm.DB) user.UserRepository {
	return &UserRepoImpl{db}
}

// Create function is use to create new user
func (c *UserRepoImpl) Create(user *model.User) (*model.User, error) {
	if err := c.db.Table("user").Save(&user).Error; err != nil {
		fmt.Errorf("Having error : %w", err)
		logrus.Error(err)
		return nil, errors.New("add user data : error")
	}
	return user, nil
}

// ReadAll function is use to get all user data
func (c *UserRepoImpl) ReadAll() ([]*model.User, error) {
	userList := make([]*model.User, 0)
	if err := c.db.Table("user").Find(&userList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("get user list data : error ")
	}
	return userList, nil
}

// Read function is get user data by id
func (c *UserRepoImpl) Read(id int) (*model.User, error) {
	user := new(model.User)

	if err := c.db.Table("user").Where("user_id = ?", id).First(&user).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("get user data : error ")
	}

	return user, nil
}

// Update function for update user data by id
func (c *UserRepoImpl) Update(id int, req *model.User) (*model.User, error) {
	user := new(model.User)

	if err := c.db.Table("user").Where("user_id = ?", id).First(&user).Update(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("update user data : error ")
	}

	return user, nil
}

// Delete function for delete user data by id
func (c *UserRepoImpl) Delete(id int) (*model.User, error) {

	user := new(model.User)

	if err := c.db.Table("user").First(&user, id).Error; err != nil {
		return nil, errors.New("id is doesnt exists")
	}

	if err := c.db.Table("user").Where("user_id = ?", id).Delete(&model.User{}).Error; err != nil {
		return nil, errors.New("delete courier data: error")
	}

	return nil, nil
}
