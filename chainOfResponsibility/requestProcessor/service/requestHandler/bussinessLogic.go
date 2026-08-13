package requesthandler

import (
	"fmt"
	"requestprocessor/middleware"
	"requestprocessor/model"
)

type BussinessLogicHandler struct {
	middleware.BaseHandler
}

func NewBussinessLogicHandler() *BussinessLogicHandler {
	return &BussinessLogicHandler{}
}

func (this BussinessLogicHandler) Handle(req model.Request) {
	fmt.Println("[BusinessLogicHandler] Processing request for", req.User, "...")
}
