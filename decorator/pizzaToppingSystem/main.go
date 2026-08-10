package main

import (
	"fmt"
	"pizzatoppingsystem/pizza"
	"pizzatoppingsystem/pizza/decorator"
)

func main() {
	pizza := decorator.NewMushroomTopping(decorator.NewCheeseTopping(pizza.NewPizza(10.0)))

	fmt.Println(pizza.GetCost())
	fmt.Println(pizza.GetDescription())
}
