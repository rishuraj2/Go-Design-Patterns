package shippingcostservice

import "shippingcalculator/order"

type FlatRateShipping struct {
	rate float64
}

func NewFlatRateShipping(rate float64) FlatRateShipping {
	return FlatRateShipping{
		rate: rate,
	}
}

func (this FlatRateShipping) CalculateCost(order order.Order) float64 {
	return this.rate
}
