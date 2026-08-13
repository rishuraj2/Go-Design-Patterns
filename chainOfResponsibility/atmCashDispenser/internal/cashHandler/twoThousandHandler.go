package cashhandler

import (
	"atmcashdispenser/model"
	"fmt"
)

type TwoThousandHandler struct {
	BaseHandler
}

func NewTwoThousandHandler() *TwoThousandHandler {
	return &TwoThousandHandler{
		BaseHandler: NewBaseHandler(2000),
	}
}

func (this *TwoThousandHandler) Handle(req *model.CashRequest) {
	if req.Amount >= this.denomination {
		noteCount := req.Amount / this.denomination
		req.Amount = req.Amount % this.denomination
		fmt.Printf("Dispensing %d x ₹%d\n", noteCount, this.denomination)
	}
	this.Forward(req)
}
