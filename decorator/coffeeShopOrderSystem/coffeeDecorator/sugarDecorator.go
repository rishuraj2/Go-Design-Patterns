package coffeedecorator

import "fmt"

type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(inner Coffee) SugarDecorator {
	return SugarDecorator{
		CoffeeDecorator: NewCoffeeDecorator(inner),
	}
} 

func (this SugarDecorator) GetCost() float64 {
	return this.inner.GetCost() + 0.20
}

func (this SugarDecorator) GetDescription() string {
	return fmt.Sprintf("%s, sugar", this.inner.GetDescription())
}
