package middleware

import (
	"fmt"
	"requestprocessor/model"
)

type AuthorizationMiddleware struct {
	BaseHandler
}

func NewAuthorizationMiddleware() *AuthorizationMiddleware {
	return &AuthorizationMiddleware{}
}

func (this *AuthorizationMiddleware) Handle(req model.Request) {
	if req.UserRole != "ADMIN" {
		fmt.Println("[AuthorizationError] invalid user")
		return
	}

	this.Forward(req)
}
