package main

import (
	"fmt"
	"shippingcalculator/order"
	shippingcostservice "shippingcalculator/shippingCostService"
)

func main() {
	ord1 := order.NewOrder("ORD-001", 10, 10)
	strategy := shippingcostservice.NewDistanceBasedShipping(40)

	sh := shippingcostservice.NewShippingCostService(strategy)
	fmt.Println(sh.CalculateShippingCost(ord1))
}
