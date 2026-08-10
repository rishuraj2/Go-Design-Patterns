package coffeedecorator

type Coffee interface {
	GetCost() float64
	GetDescription() string
}

type CoffeeDecorator struct {
	inner Coffee
}

func NewCoffeeDecorator(inner Coffee) CoffeeDecorator {
	return CoffeeDecorator{
		inner: inner,
	}
}
