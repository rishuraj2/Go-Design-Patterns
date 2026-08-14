package domain

import (
	"usermanagement/internal/repository"
	"usermanagement/model"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (this *UserService) RegisterUser(username, email string) {
	user := model.NewUser(username+email, username, email)
	this.repository.Save(user)
}

func (this *UserService) FindUser(id string) model.User {
	return this.repository.FindById(id)
}

func (this *UserService) GetAllUsers() []model.User {
	return this.repository.FindAll()
}
