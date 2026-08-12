package shippingcostservice

import "shippingcalculator/order"

type DistanceBasedShipping struct {
	ratePerKm float64
}

func NewDistanceBasedShipping(rate float64) DistanceBasedShipping {
	return DistanceBasedShipping{
		ratePerKm: rate,
	}
}

func (this DistanceBasedShipping) CalculateCost(order order.Order) float64 {
	return this.ratePerKm * order.GetDistance()
}
