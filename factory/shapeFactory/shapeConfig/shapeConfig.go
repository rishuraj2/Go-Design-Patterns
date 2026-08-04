package shapeconfig

type ShapeConfig struct {
	radius float64
	height float64
	width  float64
	base   float64
}

func (s *ShapeConfig) GetRadius() float64 {
	return s.radius
}

func (s *ShapeConfig) GetHeight() float64 {
	return s.height
}

func (s *ShapeConfig) GetWidth() float64 {
	return s.width
}

func (s *ShapeConfig) GetBase() float64 {
	return s.base
}
