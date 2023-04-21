package ray_tracer

import "testing"

func TestColor(t *testing.T) {
	r, g, b := -0.5, 0.4, 1.7

	color := NewColor(r, g, b)
	if color.R != r {
		t.Errorf("color.R = %f, got %f", r, color.R)
	}
	if color.G != g {
		t.Errorf("color.G = %f, got %f", g, color.G)
	}
	if color.B != b {
		t.Errorf("color.B = %f, got %f", b, color.B)
	}
}

func TestColorAdd(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(1.6, 0.7, 1.0)
	actual := c1.Add(c2)
	if !actual.Equals(expected) {
		t.Errorf("%+v.Add(%+v) = %f, got %f", c1, c2, expected, actual)
	}
}

func TestColorSubtract(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(0.2, 0.5, 0.5)
	actual := c1.Subtract(c2)
	if !actual.Equals(expected) {
		t.Errorf("%+v.Subtract(%+v) = %+v, got %+v", c1, c2, expected, actual)
	}
}

func TestColorMultiplyScalar(t *testing.T) {
	c1 := NewColor(0.2, 0.3, 0.4)
	factor := float64(2)
	expected := NewColor(0.4, 0.6, 0.8)
	actual := c1.MultiplyScalar(factor)
	if !actual.Equals(expected) {
		t.Errorf("%+v.MutliplyScalar(%f) = %+v, got %+v", c1, factor, expected, actual)
	}
}

func TestColorMultiply(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	expected := NewColor(0.9, 0.2, 0.04)
	actual := c1.Multiply(c2)
	if !actual.Equals(expected) {
		t.Errorf("%+v.Subtract(%+v) = %+v, got %+v", c1, c2, expected, actual)
	}
}
