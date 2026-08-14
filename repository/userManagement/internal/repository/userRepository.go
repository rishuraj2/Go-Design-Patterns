package repository

import "usermanagement/model"

type UserRepository interface {
	FindById(id string) model.User
	FindAll() []model.User
	Save(user model.User)
	Delete(id string)
}
