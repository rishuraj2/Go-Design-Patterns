package middleware

import (
	"requestprocessor/model"
)

type BaseHandler struct {
	next RequestHandler
}

func (this *BaseHandler) SetNext(next RequestHandler) {
	this.next = next
}

func (this *BaseHandler) Forward(req model.Request) {
	if this.next != nil {
		this.next.Handle(req)
	}
}
