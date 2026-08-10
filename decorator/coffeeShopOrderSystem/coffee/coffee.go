package coffee

type Coffee struct {
	price float64
}

func NewCoffee(price float64) Coffee {
	return Coffee{
		price: price,
	}
}

func (this Coffee) GetCost() float64 {
	return this.price
}

func (this Coffee) GetDescription() string {
	return "Simple Coffee"
}
