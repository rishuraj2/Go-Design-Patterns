package shape

import (
	"fmt"
	"math"
	shapeconfig "shapefactory/shapeConfig"
)

type Circle struct {
	radius float64
}

func init() {
	r := getShapeRegistry()
	err := r.register("circle", Circle{})
	if err != nil {
		fmt.Println(err)
	}
}

func (c Circle) CreateShape(config shapeconfig.ShapeConfig) Shape {
	return Circle {
		radius: config.GetRadius(),
	}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Describe() {
	fmt.Printf("Circle with radius: %f\n", c.radius)
}
