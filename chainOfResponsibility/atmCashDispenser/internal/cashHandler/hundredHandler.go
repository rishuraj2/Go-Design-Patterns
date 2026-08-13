package cashhandler

import (
	"atmcashdispenser/model"
	"fmt"
)

type HundredHandler struct {
	BaseHandler
}

func NewHundredHandler() *HundredHandler {
	return &HundredHandler{
		BaseHandler: NewBaseHandler(100),
	}
}

func (this *HundredHandler) Handle(req *model.CashRequest) {
	if req.Amount >= this.denomination {
		noteCount := req.Amount / this.denomination
		req.Amount = req.Amount % this.denomination
		fmt.Printf("Dispensing %d x ₹%d\n", noteCount, this.denomination)
	}
	this.Forward(req)
}
