package usecase

import (
	"LionChallenge/model"
	"LionChallenge/user"
	"LionChallenge/utils"
)

// UserUsecaseImpl struct
type UserUsecaseImpl struct {
	userRepo user.UserRepository
}

// CreateUserUsecaseImpl constructor implement to business logic
func CreateUserUsecaseImpl(userRepo user.UserRepository) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

// Create function is use to create new user
func (c *UserUsecaseImpl) Create(req *model.User) (map[string]interface{}, error) {
	response, err := c.userRepo.Create(req)
	if err != nil {
		return utils.Message(false, err.Error()), err
	}

	mapResponse := utils.Message(true, "create user data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}

// ReadAll bussiness logic
func (c *UserUsecaseImpl) ReadAll() (map[string]interface{}, error) {
	response, err := c.userRepo.ReadAll()
	if err != nil {
		return utils.Message(false, err.Error()), err
	}

	mapResponse := utils.Message(true, "read all user data: success ")
	mapResponse["data"] = response
	return mapResponse, nil
}

func (c *UserUsecaseImpl) Read(id int) (map[string]interface{}, error) {
	response, err := c.userRepo.Read(id)
	if err != nil {
		return utils.Message(false, err.Error()), err
	}

	mapResponse := utils.Message(true, "read user data: success ")
	mapResponse["data"] = response
	return mapResponse, nil
}

// Update function is use to update user data by id
func (c *UserUsecaseImpl) Update(id int, req *model.User) (map[string]interface{}, error) {
	response, err := c.userRepo.Update(id, req)
	if err != nil {
		return utils.Message(false, err.Error()), err
	}

	mapResponse := utils.Message(true, "update user data: success ")
	mapResponse["response"] = response
	return mapResponse, nil
}

// Delete function
func (c *UserUsecaseImpl) Delete(id int) (map[string]interface{}, error) {
	_, err := c.userRepo.Delete(id)
	if err != nil {
		return utils.Message(false, err.Error()), err
	}

	mapResponse := utils.Message(true, "delete user data: success ")
	return mapResponse, nil

}
