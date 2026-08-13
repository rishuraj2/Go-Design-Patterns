package cashhandler

import (
	"atmcashdispenser/model"
	"fmt"
)

type OnesHandler struct {
	BaseHandler
}

func NewOnesHandler() *OnesHandler {
	return &OnesHandler{
		BaseHandler: NewBaseHandler(1),
	}
}

func (this *OnesHandler) Handle(req *model.CashRequest) {
	if req.Amount >= this.denomination {
		noteCount := req.Amount / this.denomination
		req.Amount = req.Amount % this.denomination
		fmt.Printf("Dispensing %d x ₹%d\n", noteCount, this.denomination)
	}
	this.Forward(req)
}
