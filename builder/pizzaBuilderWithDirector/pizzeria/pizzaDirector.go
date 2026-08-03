package pizzeria

type PizzaDirector struct{}

func NewPizzaDirector() PizzaDirector {
	return PizzaDirector{}
}

func (pd PizzaDirector) BuildMargherita(size pizzaSize) *pizza {
	pizza := NewPizzaBuilder(size).
		Crust(THIN).
		Sauce(MARINARA).
		Cheese(MOZZARELLA).
		Build()

	return pizza
}

func (pd PizzaDirector) BuildPepperoni(size pizzaSize) *pizza {
	pizza := NewPizzaBuilder(size).
		Crust(THICK).
		Sauce(MARINARA, PESTO).
		Cheese(MOZZARELLA, CHEDDAR).
		Toppings(PEPPERONI).
		Build()

	return pizza
}

func (pd PizzaDirector) BuildVeggie(size pizzaSize) *pizza {
	pizza := NewPizzaBuilder(size).
		Crust(THIN).
		Sauce(MARINARA).
		Cheese(VEGAN).
		Toppings(MUSHROOMS, ONIONS, OLIVES).
		Build()

	return pizza
}
