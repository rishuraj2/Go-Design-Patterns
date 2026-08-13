package cashhandler

import "atmcashdispenser/model"

type Handler interface {
	Handle(req *model.CashRequest)
	SetNext(next Handler)
}

type BaseHandler struct {
	next         Handler
	denomination int
}

func NewBaseHandler(denomination int) BaseHandler {
	return BaseHandler{
		denomination: denomination,
	}
}

func (this *BaseHandler) SetNext(handler Handler) {
	this.next = handler
}

func (this *BaseHandler) Forward(req *model.CashRequest) {
	if this.next != nil {
		this.next.Handle(req)
	}
}
