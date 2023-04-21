package components

type Point struct {
	X, Y, Z float64
}

func NewPoint(X, Y, Z float64) Point {
	return Point{X, Y, Z}
}

func (p Point) AddVector(v Vector) Point {
	return NewPoint(p.X+v.X, p.Y+v.Y, p.Z+v.Z)
}

func (p Point) Subtract(p1 Point) Vector {
	return NewVector(p.X-p1.X, p.Y-p1.Y, p.Z-p1.Z)
}

func (p Point) SubtractVector(v Vector) Point {
	return NewPoint(p.X-v.X, p.Y-v.Y, p.Z-v.Z)
}

func (p Point) Equals(p1 Point) bool {
	return AlmostEqual(p.X, p1.X) && AlmostEqual(p.Y, p1.Y) && AlmostEqual(p.Z, p1.Z)
}
