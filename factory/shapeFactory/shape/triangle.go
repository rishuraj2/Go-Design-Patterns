package shape

import (
	"fmt"
	shapeconfig "shapefactory/shapeConfig"
)

type Triangle struct {
	base   float64
	height float64
}

func init() {
	r := getShapeRegistry()
	err := r.register("triangle", Triangle{})
	if err != nil {
		fmt.Println(err)
	}
}

func (t Triangle) CreateShape(config shapeconfig.ShapeConfig) Shape {
	return Triangle{
		base: config.GetBase(),
		height: config.GetHeight(),
	}
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Describe() {
	fmt.Printf("Triangle with base: %f and height: %f", t.base, t.height)
}
