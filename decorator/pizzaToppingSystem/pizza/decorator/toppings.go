package decorator

type Pizza interface {
	GetCost() float64
	GetDescription() string
}

type Toppings struct {
	inner Pizza
}

func NewToppings(inner Pizza) Toppings {
	return Toppings{
		inner: inner,
	}
}
