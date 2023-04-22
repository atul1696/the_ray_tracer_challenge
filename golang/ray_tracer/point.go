package ray_tracer

import (
	"fmt"
)

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
	return almostEqual(p.X, p1.X) && almostEqual(p.Y, p1.Y) && almostEqual(p.Z, p1.Z)
}

func (p Point) Tuple() Matrix {
	tp := NewTuple(4)
	tp[0][0] = p.X
	tp[1][0] = p.Y
	tp[2][0] = p.Z
	tp[3][0] = 1.0
	return tp
}

func (p Point) Transform(m Matrix) (Point, error) {
	t := p.Tuple()
	transform, err := m.Multiply(t)
	if err != nil {
		return Point{}, fmt.Errorf("cannot apply transformation to point, uncompatible matrix [%d, %d]", m.Rows(), m.Columns())
	}
	p1, _ := transform.Point()
	return p1, nil
}
