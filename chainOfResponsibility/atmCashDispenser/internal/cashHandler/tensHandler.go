package cashhandler

import (
	"atmcashdispenser/model"
	"fmt"
)

type TensHandler struct {
	BaseHandler
}

func NewTensHandler() *TensHandler {
	return &TensHandler{
		BaseHandler: NewBaseHandler(10),
	}
}

func (this *TensHandler) Handle(req *model.CashRequest) {
	if req.Amount >= this.denomination {
		noteCount := req.Amount / this.denomination
		req.Amount = req.Amount % this.denomination
		fmt.Printf("Dispensing %d x ₹%d\n", noteCount, this.denomination)
	}
	this.Forward(req)
}
