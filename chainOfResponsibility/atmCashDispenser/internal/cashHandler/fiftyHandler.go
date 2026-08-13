package cashhandler

import (
	"atmcashdispenser/model"
	"fmt"
)

type FiftyHandler struct {
	BaseHandler
}

func NewFiftyHandler() *FiftyHandler {
	return &FiftyHandler{
		BaseHandler: NewBaseHandler(50),
	}
}

func (this *FiftyHandler) Handle(req *model.CashRequest) {
	if req.Amount >= this.denomination {
		noteCount := req.Amount / this.denomination
		req.Amount = req.Amount % this.denomination
		fmt.Printf("Dispensing %d x ₹%d\n", noteCount, this.denomination)
	}
	this.Forward(req)
}
