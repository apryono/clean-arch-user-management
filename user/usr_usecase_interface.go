package user

import "LionChallenge/model"

// UserUsecase interface
type UserUsecase interface {
	Create(req *model.User) (map[string]interface{}, error)
	Read(id int) (map[string]interface{}, error)
	ReadAll() (map[string]interface{}, error)
	Update(id int, req *model.User) (map[string]interface{}, error)
	Delete(id int) (map[string]interface{}, error)
}
