package ray_tracer

import (
	"testing"
)

func TestMaterialLighting(t *testing.T) {
	white := NewColor(1, 1, 1)
	tests := []struct{
		ev Vector
		l PointLight
		expected Color
		inShadow bool
	}{
		{NewVector(0, 0, -1), NewPointLight(NewPoint(0, 0, -10), white), NewColor(1.9, 1.9, 1.9), false},
		{NewVector(0, 0.707107, -0.707107), NewPointLight(NewPoint(0, 0, -10), white), NewColor(1.0, 1.0, 1.0), false},
		{NewVector(0, 0, -1), NewPointLight(NewPoint(0, 10, -10), white), NewColor(0.7364, 0.7364, 0.7364), false},
		{NewVector(0, -0.707107, -0.707107), NewPointLight(NewPoint(0, 10, -10), white), NewColor(1.63639, 1.63639, 1.63639), false},
		{NewVector(0, 0, -1), NewPointLight(NewPoint(0, 0, 10), white), NewColor(0.1, 0.1, 0.1), false},
		{NewVector(0, 0, -1), NewPointLight(NewPoint(0, 0, -10), white), NewColor(0.1, 0.1, 0.1), true},
	}
	nv := NewVector(0, 0, -1)
	m := NewMaterial()
	p := NewPoint(0, 0, 0)
	for _, test := range tests {
		actual := Lighting(m, p, test.l, test.ev, nv, test.inShadow)
		if !test.expected.Equals(actual) {
			t.Errorf("for eye: %+v, light: %+v expected %+v got %+v", test.ev, test.l, test.expected, actual)
		}
	}
}
