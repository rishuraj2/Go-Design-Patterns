package order

type Order struct {
	id       string
	weight   float64
	distance float64
}

func NewOrder(id string, weight, distance float64) Order {
	return Order{
		id:       id,
		weight:   weight,
		distance: distance,
	}
}

func (this Order) GetWeight() float64 {
	return this.weight
}

func (this Order) GetDistance() float64 {
	return this.distance
}
