package shape

import shapeconfig "shapefactory/shapeConfig"

type Shape interface {
	Area() float64
	Describe()
	CreateShape(config shapeconfig.ShapeConfig) Shape
}
