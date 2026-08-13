package main

import (
	cashhandler "atmcashdispenser/internal/cashHandler"
	chainbuilder "atmcashdispenser/internal/chainBuilder"
	"atmcashdispenser/model"
)

func main() {
	twoThousand := cashhandler.NewTwoThousandHandler()
	fiveHundred := cashhandler.NewFiveHundredHandler()
	oneHundred := cashhandler.NewHundredHandler()
	fifty := cashhandler.NewFiftyHandler()
	tens := cashhandler.NewTensHandler()
	ones := cashhandler.NewOnesHandler()

	builder := chainbuilder.NewChainBuilder(twoThousand).
	Chain(fiveHundred).
	Chain(oneHundred).
	Chain(fifty).
	Chain(tens).
	Chain(ones).
	Build()

	req := &model.CashRequest{
		Amount: 45677,
	}

	builder.Handle(req)
}