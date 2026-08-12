package shippingcostservice

import "shippingcalculator/order"

type ShippingStrategy interface {
	CalculateCost(order order.Order) float64
}

type ShippingCostService struct {
	strategy ShippingStrategy
}

func NewShippingCostService(stategy ShippingStrategy) ShippingCostService {
	return ShippingCostService{
		strategy: stategy,
	}
}

func (this *ShippingCostService) SetStrategy(strategy ShippingStrategy) {
	this.strategy = strategy
}

func (this ShippingCostService) CalculateShippingCost(order order.Order) float64 {
	return this.strategy.CalculateCost(order)
}
