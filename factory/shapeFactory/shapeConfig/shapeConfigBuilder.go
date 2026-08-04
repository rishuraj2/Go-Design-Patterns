package shapeconfig

type ShapeConfigBuilder struct {
	radius float64
	height float64
	width  float64
	base   float64
}

func NewShapeConfigBuilder() *ShapeConfigBuilder {
	return &ShapeConfigBuilder{}
}

func (sb *ShapeConfigBuilder) Radius(radius float64) *ShapeConfigBuilder {
	sb.radius = radius
	return sb
}

func (sb *ShapeConfigBuilder) Height(height float64) *ShapeConfigBuilder {
	sb.height = height
	return sb
}

func (sb *ShapeConfigBuilder) Width(width float64) *ShapeConfigBuilder {
	sb.width = width
	return sb
}

func (sb *ShapeConfigBuilder) Base(base float64) *ShapeConfigBuilder {
	sb.base = base
	return sb
}

func (sb *ShapeConfigBuilder) Build() ShapeConfig {
	return ShapeConfig{
		radius: sb.radius,
		height: sb.height,
		width:  sb.width,
		base:   sb.base,
	}
}
