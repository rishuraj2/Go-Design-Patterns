package main

import "pizzabuilder/pizzeria"

func main() {
	custom := pizzeria.NewPizzaBuilder(pizzeria.MEDIUM).Crust(pizzeria.STUFFED).Sauce(pizzeria.MARINARA).Cheese(pizzeria.CHEDDAR, pizzeria.GORGONZOLA).Toppings(pizzeria.MUSHROOMS, pizzeria.OLIVES).Build()
	custom.Summary()
}
