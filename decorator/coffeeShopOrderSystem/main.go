package main

import (
	"coffeeshopordersystem/coffee"
	coffeedecorator "coffeeshopordersystem/coffeeDecorator"
	"fmt"
)

func main() {
	ord1 := coffeedecorator.NewSugarDecorator(coffeedecorator.NewMilkDecorator(coffee.NewCoffee(10.00)))
	fmt.Println(ord1.GetDescription())
	fmt.Println(ord1.GetCost())
}
