package decorator

import "fmt"

type CheeseTopping struct {
	Toppings
}

func NewCheeseTopping(inner Pizza) CheeseTopping {
	return CheeseTopping{
		Toppings: NewToppings(inner),
	}
}

func (this CheeseTopping) GetCost() float64 {
	return this.inner.GetCost() + 2.25
}

func (this CheeseTopping) GetDescription() string {
	return fmt.Sprintf("%s, cheese", this.inner.GetDescription())
}