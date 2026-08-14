package database

import (
	"sync"
	"usermanagement/model"
)

type InMemoryStore struct {
	store map[string]model.User
}

var (
	instance *InMemoryStore
	once     sync.Once
)

func NewInMemoryStore() *InMemoryStore {
	once.Do(func() {
		instance = &InMemoryStore{
			store: make(map[string]model.User),
		}
	})

	return instance
}

func (this *InMemoryStore) FindById(id string) model.User {
	if user, exists := this.store[id]; exists {
		return user
	}

	return model.User{}
}

func (this *InMemoryStore) FindAll() []model.User {
	var users []model.User

	for _, user := range this.store {
		users = append(users, user)
	}

	return users
}

func (this *InMemoryStore) Save(user model.User) {
	this.store[user.GetId()] = user
}

func (this *InMemoryStore) Delete(id string) {
	delete(this.store, id)
}
