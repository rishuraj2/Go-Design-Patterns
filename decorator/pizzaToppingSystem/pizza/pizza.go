package pizza

type Pizza struct {
	price float64
}

func NewPizza(price float64) Pizza {
	return Pizza{
		price: price,
	}
}

func (this Pizza) GetCost() float64 {
	return this.price
}

func (this Pizza) GetDescription() string {
	return "Pizza"
}
