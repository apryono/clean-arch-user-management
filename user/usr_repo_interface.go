package user

import "LionChallenge/model"

// UserRepository interface
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Read(id int) (*model.User, error)
	ReadAll() ([]*model.User, error)
	Update(id int, req *model.User) (*model.User, error)
	Delete(id int) (*model.User, error)
}
