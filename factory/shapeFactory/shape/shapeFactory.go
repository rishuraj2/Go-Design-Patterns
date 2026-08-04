package shape

import shapeconfig "shapefactory/shapeConfig"

type ShapeFactory struct {
	registry *shapeRegistry
}

func NewShapeFactory() ShapeFactory {
	return ShapeFactory{
		registry: getShapeRegistry(),
	}
}

func (fact ShapeFactory) CreateShape(name string, config shapeconfig.ShapeConfig) (Shape, error) {
	sh, err := fact.registry.getShape(name)

	if err != nil {
		return nil, err
	}

	return sh.CreateShape(config), nil
}
