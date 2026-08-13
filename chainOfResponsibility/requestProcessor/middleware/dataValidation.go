package middleware

import (
	"fmt"
	"requestprocessor/model"
	"strings"
)

type DataValidationMiddleware struct {
	BaseHandler
}

func NewDataValidationMiddleware() *DataValidationMiddleware {
	return &DataValidationMiddleware{}
}

func (this *DataValidationMiddleware) Handle(req model.Request) {
	if strings.TrimSpace(req.Payload) == "" {
		fmt.Println("[Data Validation Error] Invalid Payload")
		return
	}

	this.Forward(req)
}
