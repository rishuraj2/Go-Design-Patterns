package pizzeria

type PizzaBuilder struct {
	size     pizzaSize
	crust    pizzaCrust
	sauces   pizzaSauces
	cheeses  pizzaCheeses
	toppings pizzaToppings
}

func NewPizzaBuilder(size pizzaSize) *PizzaBuilder {
	return &PizzaBuilder{
		size: size,
	}
}

func (pb *PizzaBuilder) Crust(crust pizzaCrust) *PizzaBuilder {
	pb.crust = crust
	return pb
}

func (pb *PizzaBuilder) Sauce(sauces ...pizzaSauce) *PizzaBuilder {
	for _, sauce := range sauces {
		pb.sauces = append(pb.sauces, sauce)
	}
	return pb
}

func (pb *PizzaBuilder) Cheese(cheeses ...pizzaCheese) *PizzaBuilder {
	for _, cheese := range cheeses {
		pb.cheeses = append(pb.cheeses, cheese)
	}
	return pb
}

func (pb *PizzaBuilder) Toppings(toppings ...pizzaTopping) *PizzaBuilder {
	for _, topping := range toppings {
		pb.toppings = append(pb.toppings, topping)
	}
	return pb
}

func (pb *PizzaBuilder) Build() *pizza {
	return &pizza{
		size:     pb.size,
		crust:    pb.crust,
		sauce:    pb.sauces,
		cheese:   pb.cheeses,
		toppings: pb.toppings,
	}
}
