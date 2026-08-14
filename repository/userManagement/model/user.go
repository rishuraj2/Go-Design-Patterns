package model

type User struct {
	id       string
	username string
	email    string
}

func NewUser(id, username, email string) User {
	return User{
		id:       id,
		username: username,
		email:    email,
	}
}

func (this User) GetId() string {
	return this.id
}

func (this User) GetUsername() string {
	return this.username
}

func (this User) GetEmail() string {
	return this.email
}
