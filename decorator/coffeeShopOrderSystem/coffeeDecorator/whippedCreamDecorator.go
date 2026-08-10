package coffeedecorator

import "fmt"

type WhippedCreamDecorator struct {
	CoffeeDecorator
}

func NewWhippedCreamDecorator(inner Coffee) WhippedCreamDecorator {
	return WhippedCreamDecorator{
		CoffeeDecorator: NewCoffeeDecorator(inner),
	}
}

func (this WhippedCreamDecorator) GetCost() float64 {
	return this.inner.GetCost() + 1.00
}

func (this WhippedCreamDecorator) GetDescription() string {
	return fmt.Sprintf("%s, whipperCream", this.inner.GetDescription())
}
