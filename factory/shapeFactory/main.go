package main

import (
	"fmt"
	"shapefactory/shape"
	shapeconfig "shapefactory/shapeConfig"
)

func main() {
	conf := shapeconfig.NewShapeConfigBuilder().
		Height(10).
		Width(10).
		Build()

	shape, err := shape.NewShapeFactory().CreateShape("rectangle", conf)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(shape.Area())
	shape.Describe()
}
