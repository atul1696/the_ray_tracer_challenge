package components

import (
	"testing"
)

func TestPoint(t *testing.T) {
	x, y, z := 4.3, -4.2, 3.1
	
	point := NewPoint(x, y, z)
	if point.X != x { 
		t.Errorf("tuple.X = %f, got %f", x, point.X)
	}
	if point.Y != y { 
		t.Errorf("tuple.Y = %f, got %f", y, point.Y)
	}
	if point.Z != z { 
		t.Errorf("tuple.Z = %f, got %f", z, point.Z)
	}
}

func TestPointAddVector(t *testing.T) {
	p := NewPoint(3, -2, 5)
	v := NewVector(-2, 3, 1)
	actual := p.AddVector(v)
	expected := NewPoint(1, 1, 6)
	if actual != expected {
		t.Errorf("%+v.AddVector(%+v) = %+v, got %+v", p, v, expected, actual)
	}
}

func TestPointSubtractPoint(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)
	actual := p1.SubtractPoint(p2)
	expected := NewVector(-2, -4, -6)
	if actual != expected {
		t.Errorf("%+v.SubtractPoint(%+v) = %+v, got %+v", p1, p2, expected, actual)
	}
}

func TestPointSubtractVector(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v := NewVector(5, 6, 7)
	actual := p.SubtractVector(v)
	expected := NewPoint(-2, -4, -6)
	if actual != expected {
		t.Errorf("%+v.SubtractVector(%+v) = %+v, got %+v", p, v, expected, actual)
	}
}
