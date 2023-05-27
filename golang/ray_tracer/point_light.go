package ray_tracer

type PointLight struct {
	Position  Point
	Intensity Color
}

func NewPointLight(p Point, i Color) PointLight {
	return PointLight{p, i}
}
