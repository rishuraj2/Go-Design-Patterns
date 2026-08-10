package coffeedecorator

import "fmt"

type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(inner Coffee) MilkDecorator {
	return MilkDecorator{
		CoffeeDecorator: NewCoffeeDecorator(inner),
	}
}

func (this MilkDecorator) GetCost() float64 {
	return this.inner.GetCost() + 0.50
}

func (this MilkDecorator) GetDescription() string {
	return fmt.Sprintf("%s, milk", this.inner.GetDescription())
}
