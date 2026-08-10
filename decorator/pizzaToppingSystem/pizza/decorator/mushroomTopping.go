package decorator

import "fmt"

type MushroomTopping struct {
	Toppings
}

func NewMushroomTopping(inner Pizza) MushroomTopping {
	return MushroomTopping{
		Toppings: NewToppings(inner),
	}
}

func (this MushroomTopping) GetCost() float64 {
	return this.inner.GetCost() + 5.0
}

func (this MushroomTopping) GetDescription() string {
	return fmt.Sprintf("%s, mushroom", this.inner.GetDescription())
}
