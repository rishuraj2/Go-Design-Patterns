package shippingcostservice

import "shippingcalculator/order"

type WeightBasedShipping struct {
	ratePerKg float64
}

func NewWeightBasedShipping(rate float64) WeightBasedShipping {
	return WeightBasedShipping{
		ratePerKg: rate,
	}
}

func (this WeightBasedShipping) CalculateCost(order order.Order) float64 {
	return this.ratePerKg * order.GetWeight()
}
