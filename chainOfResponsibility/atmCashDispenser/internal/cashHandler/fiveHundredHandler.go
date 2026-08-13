package cashhandler

import (
	"atmcashdispenser/model"
	"fmt"
)

type FiveHundredHandler struct {
	BaseHandler
}

func NewFiveHundredHandler() *FiveHundredHandler {
	return &FiveHundredHandler{
		BaseHandler: NewBaseHandler(500),
	}
}

func (this *FiveHundredHandler) Handle(req *model.CashRequest) {
	if req.Amount >= this.denomination {
		noteCount := req.Amount / this.denomination
		req.Amount = req.Amount % this.denomination
		fmt.Printf("Dispensing %d x ₹%d\n", noteCount, this.denomination)
	}
	this.Forward(req)
}
