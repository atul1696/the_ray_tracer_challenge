package components

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func NewVector(X, Y, Z float64) Vector {
	return Vector{X, Y, Z}
}

func (v Vector) Equals(v1 Vector) bool {
	return AlmostEqual(v.X, v1.X) && AlmostEqual(v.Y, v1.Y) && AlmostEqual(v.Z, v1.Z)
}

func (v Vector) AddVector(v1 Vector) Vector {
	return NewVector(v.X + v1.X, v.Y + v1.Y, v.Z + v1.Z)
}

func (v Vector) SubtractVector(v1 Vector) Vector {
	return NewVector(v.X - v1.X, v.Y - v1.Y, v.Z - v1.Z)
}

func (v Vector) Negate() Vector {
	return NewVector(-v.X, -v.Y, -v.Z)
}

func (v Vector) MultiplyScalar(factor float64) Vector {
	return NewVector(v.X * factor, v.Y * factor, v.Z * factor)
}

func (v Vector) DivideScalar(factor float64) Vector {
	return v.MultiplyScalar(1 / factor)
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	return NewVector(v.X / mag, v.Y / mag, v.Z / mag)
}

func DotProduct(v1, v2 Vector) float64 {
	return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func CrossProduct(v1, v2 Vector) Vector {
	return NewVector(
		v1.Y * v2.Z - v1.Z * v2.Y,
		v1.Z * v2.X - v1.X * v2.Z,
		v1.X * v2.Y - v1.Y * v2.X,
	)
}