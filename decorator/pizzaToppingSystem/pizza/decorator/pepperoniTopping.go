package decorator

import "fmt"

type PepperoniTopping struct {
	Toppings
}

func NewPepperoniTopping(inner Pizza) PepperoniTopping {
	return PepperoniTopping{
		Toppings: NewToppings(inner),
	}
}

func (this PepperoniTopping) GetCost() float64 {
	return this.inner.GetCost() + 10.0
}

func (this PepperoniTopping) GetDescription() string {
	return fmt.Sprintf("%s, pepperoni", this.inner.GetDescription())
}
