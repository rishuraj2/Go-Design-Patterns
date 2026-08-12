package shippingcostservice

import "shippingcalculator/order"

type ThirdPartyApiShipping struct {
	baseFee       float64
	percentageFee float64
}

func NewThirdPartyApiShipping(baseFee, percentageFee float64) ThirdPartyApiShipping {
	return ThirdPartyApiShipping{
		baseFee:       baseFee,
		percentageFee: percentageFee,
	}
}

func (this ThirdPartyApiShipping) CalculateCost(order order.Order) float64 {
	return this.baseFee + this.percentageFee
}
