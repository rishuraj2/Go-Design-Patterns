package middleware

import (
	"fmt"
	"requestprocessor/model"
)

type AuthenticationMiddleware struct {
	BaseHandler
}

func NewAuthenticationMiddleware() *AuthenticationMiddleware {
	return &AuthenticationMiddleware{}
}

func (this *AuthenticationMiddleware) Handle(req model.Request) {
	if req.User != "abc" {
		fmt.Println("[AuthenticationError] invalid user")
		return
	}

	this.Forward(req)
}
