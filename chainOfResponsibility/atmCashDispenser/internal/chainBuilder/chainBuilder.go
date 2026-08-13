package chainbuilder

import cashhandler "atmcashdispenser/internal/cashHandler"

type ChainBuilder struct {
	start cashhandler.Handler
	end   cashhandler.Handler
}

func NewChainBuilder(start cashhandler.Handler) *ChainBuilder {
	return &ChainBuilder{
		start: start,
		end: start,
	}
} 

func (this *ChainBuilder) Chain(handler cashhandler.Handler) *ChainBuilder {
	this.end.SetNext(handler)
	this.end = handler
	return this
}

func (this *ChainBuilder) Build() cashhandler.Handler {
	return this.start
}
