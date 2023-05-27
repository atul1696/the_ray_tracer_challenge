package ray_tracer

import (
	"math"
	"testing"
)

func TestSphereNormal(t *testing.T) {
	r3f := 1 / math.Sqrt(3)
	r2f := 1 / math.Sqrt(2)
	identity := NewIdentityMatrix(4)
	tf, _ := Scale(1, 0.5, 1).Multiply(RotateZ(math.Pi / 5))
	tests := []struct {
		p         Point
		transform Matrix
		expected  Vector
	}{
		{NewPoint(1, 0, 0), identity, NewVector(1, 0, 0)},
		{NewPoint(0, 1, 0), identity, NewVector(0, 1, 0)},
		{NewPoint(0, 0, 1), identity, NewVector(0, 0, 1)},
		{NewPoint(r3f, r3f, r3f), identity, NewVector(r3f, r3f, r3f)},
		{NewPoint(0, 1.70711, -0.70711), Translate(0, 1, 0), NewVector(0, 0.70711, -0.70711)},
		{NewPoint(0, r2f, -r2f), tf, NewVector(0, 0.97014, -0.24254)},
	}
	s := NewSphere()
	for _, test := range tests {
		s.SetTransform(test.transform)
		actual := s.NormatAt(test.p)
		if !test.expected.Equals(actual) {
			t.Errorf("at point %+v expected %+v, got %+v", test.p, test.expected, actual)
		}
	}
}
