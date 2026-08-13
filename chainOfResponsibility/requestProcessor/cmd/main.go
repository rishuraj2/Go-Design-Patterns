package main

import (
	"requestprocessor/middleware"
	"requestprocessor/model"
	requesthandler "requestprocessor/service/requestHandler"
)

func main() {
	authentication := middleware.NewAuthenticationMiddleware()
	authorization := middleware.NewAuthorizationMiddleware()
	dataValidation := middleware.NewDataValidationMiddleware()
	bussinessLogic := requesthandler.NewBussinessLogicHandler()

	authentication.SetNext(authorization)
	authorization.SetNext(dataValidation)
	dataValidation.SetNext(bussinessLogic)

	req := model.Request{
		User: "abc",
		UserRole: "ADMIN",
		RequestCount: 1000,
		Payload: "Hi!",
	}

	authentication.Handle(req)

}