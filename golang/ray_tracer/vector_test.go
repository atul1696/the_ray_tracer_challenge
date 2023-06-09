package ray_tracer

import (
	"testing"
)

func TestVector(t *testing.T) {
	x, y, z := 4.3, -4.2, 3.1

	vector := NewVector(x, y, z)
	if vector.X != x {
		t.Errorf("vector.X = %f, got %f", x, vector.X)
	}
	if vector.Y != y {
		t.Errorf("vector.Y = %f, got %f", y, vector.Y)
	}
	if vector.Z != z {
		t.Errorf("vector.Z = %f, got %f", z, vector.Z)
	}
}

func TestVectorAdd(t *testing.T) {
	v1 := NewVector(3, -2, 5)
	v2 := NewVector(-2, 3, 1)
	actual := v1.Add(v2)
	expected := NewVector(1, 1, 6)
	if actual != expected {
		t.Errorf("%+v.Add(%+v) = %+v, got %+v", v1, v2, expected, actual)
	}
}

func TestVectorSubtract(t *testing.T) {
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)
	actual := v1.Subtract(v2)
	expected := NewVector(-2, -4, -6)
	if actual != expected {
		t.Errorf("%+v.Subtract(%+v) = %+v, got %+v", v1, v2, expected, actual)
	}
}

func TestVectorNegate(t *testing.T) {
	v := NewVector(1, -2, 3)
	actual := v.Negate()
	expected := NewVector(-1, 2, -3)
	if actual != expected {
		t.Errorf("%+v.Negate() = %+v, got %+v", v, expected, actual)
	}
}

func TestVectorMutiplyScalar(t *testing.T) {
	v := NewVector(1, -2, 3)
	factor := 3.5
	actual := v.MultiplyScalar(factor)
	expected := NewVector(3.5, -7, 10.5)
	if actual != expected {
		t.Errorf("%+v.MultiplyScalar(%f) = %+v, got %+v", v, factor, expected, actual)
	}
}

func TestVectorDivideScalar(t *testing.T) {
	v := NewVector(1, -2, 3)
	factor := float64(2)
	actual := v.DivideScalar(factor)
	expected := NewVector(0.5, -1, 1.5)
	if actual != expected {
		t.Errorf("%+v.DivideScalar(%f) = %+v, got %+v", v, factor, expected, actual)
	}
}

func TestMagnitude(t *testing.T) {
	var tests = []struct {
		x        float64
		y        float64
		z        float64
		expected float64
	}{
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
		{1, 2, 3, 3.7416573868},
		{-1, -2, -3, 3.7416573868},
	}
	for _, test := range tests {
		v := NewVector(test.x, test.y, test.z)
		if actual := v.Magnitude(); !almostEqual(test.expected, actual) {
			t.Errorf("%+v.Magnitude() = %+v, got %+v", v, test.expected, actual)
		}
	}
}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		v        Vector
		expected Vector
	}{
		{NewVector(4, 0, 0), NewVector(1, 0, 0)},
		{NewVector(1, 2, 3), NewVector(0.26726, 0.53452, 0.80178)},
	}
	for _, test := range tests {
		actual := test.v.Normalize()
		if !actual.Equals(test.expected) {
			t.Errorf("%+v.Normalize() = %+v, got %+v", test.v, test.expected, actual)
		}
		mag := actual.Magnitude()
		if !almostEqual(1, mag) {
			t.Errorf("%+v.Magnitude() = %+v, got %+v", actual, 1, mag)
		}
	}
}

func TestDotProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)
	actual := DotProduct(v1, v2)
	expected := float64(20)
	if !(almostEqual(actual, expected)) {
		t.Errorf("DotProduct(%+v, %+v) = %f, got %f", v1, v2, expected, actual)
	}
}

func TestCrossProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	actual := CrossProduct(v1, v2)
	expected := NewVector(-1, 2, -1)
	if !expected.Equals(actual) {
		t.Errorf("CrossProduct(%+v, %+v) = %f, got %f", v1, v2, expected, actual)
	}

	actual = CrossProduct(v2, v1)
	expected = NewVector(1, -2, 1)
	if !expected.Equals(actual) {
		t.Errorf("CrossProduct(%+v, %+v) = %f, got %f", v1, v2, expected, actual)
	}
}

func TestVectorReflect(t *testing.T) {
	tests := []struct {
		v        Vector
		n        Vector
		expected Vector
	}{
		{NewVector(1, -1, 0), NewVector(0, 1, 0), NewVector(1, 1, 0)},
		{NewVector(0, -1, 0), NewVector(0.707107, 0.707107, 0), NewVector(1, 0, 0)},
	}
	for _, test := range tests {
		actual := test.v.Reflect(test.n)
		if !test.expected.Equals(actual) {
			t.Errorf("vector %+v reflect using normal %+v expected %+v got %+v", test.v, test.n, test.expected, actual)
		}
	}
}
