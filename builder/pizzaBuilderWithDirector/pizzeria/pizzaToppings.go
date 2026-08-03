package pizzeria

import "strings"

type pizzaTopping int

const (
	PEPPERONI pizzaTopping = iota
	MUSHROOMS
	ONIONS
	OLIVES
	PEPPERS
	SAUSAGE
	BACON
	PINEAPPLE
)

func (p pizzaTopping) String() string {
	val := []string{"pepperoni", "mushrooms", "onions", "olives", "peppers", "sausage", "bacon", "pineapple"}

	if int(p) < len(val) {
		return val[int(p)]
	}

	return "unknown"
}

type pizzaToppings []pizzaTopping

func (p pizzaToppings) String() string {
	var sb strings.Builder

	for _, topping := range p {
		sb.WriteString(topping.String() + " ")
	}

	return sb.String()
}
