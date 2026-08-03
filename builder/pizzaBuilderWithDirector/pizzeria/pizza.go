package pizzeria

import "fmt"

type pizza struct {
	size     pizzaSize
	crust    pizzaCrust
	sauce    pizzaSauces
	cheese   pizzaCheeses
	toppings pizzaToppings
}

func (p pizza) Summary() {
	fmt.Printf("Size: %s\nCrust: %s\nSauce: %s\nCheese: %s\nToppings: %s\n", p.size.String(), p.crust.String(), p.sauce.String(), p.cheese.String(), p.toppings.String())
}
