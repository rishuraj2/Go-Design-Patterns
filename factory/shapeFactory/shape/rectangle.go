package shape

import (
	"fmt"
	shapeconfig "shapefactory/shapeConfig"
)

type Rectangle struct {
	height float64
	width  float64
}

func init() {
	r := getShapeRegistry()
	err := r.register("rectangle", Rectangle{})
	if err != nil {
		fmt.Println(err)
	}
}

func (r Rectangle) CreateShape(config shapeconfig.ShapeConfig) Shape {
	return Rectangle{
		height: config.GetHeight(),
		width:  config.GetWidth(),
	}
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Describe() {
	fmt.Printf("Rectangle with height: %f and width: %f\n", r.height, r.width)
}
