package ray_tracer

import (
	"testing"
)

func TestRayCreate(t *testing.T) {
	o := NewPoint(1, 2, 3)
	d := NewVector(4, 5, 6)

	ray := NewRay(o, d)
	if !o.Equals(ray.Origin) {
		t.Errorf("ray origin %+v is not the same as %+v", ray.Origin, o)
	}
	if !d.Equals(ray.Direction) {
		t.Errorf("ray direction %+v is not the same as %+v", ray.Direction, d)
	}
}

func TestRayPosition(t *testing.T) {
	tests := []struct {
		t        float64
		expected Point
	}{
		{0, NewPoint(2, 3, 4)},
		{1, NewPoint(3, 3, 4)},
		{-1, NewPoint(1, 3, 4)},
		{2.5, NewPoint(4.5, 3, 4)},
	}
	r := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))
	for _, test := range tests {
		if actual := r.Position(test.t); !test.expected.Equals(actual) {
			t.Errorf("ray at position %f = %+v, got %+v", test.t, test.expected, actual)
		}
	}
}

func TestRayIntersect(t *testing.T) {
	tests := []struct {
		origin   Point
		expected []float64
	}{
		{NewPoint(0, 0, -5), []float64{4.0, 6.0}},
		{NewPoint(0, 1, -5), []float64{5.0, 5.0}},
		{NewPoint(0, 2, -5), []float64{}},
		{NewPoint(0, 0, 0), []float64{-1.0, 1.0}},
		{NewPoint(0, 0, 5), []float64{-6.0, -4.0}},
	}

	direction := NewVector(0, 0, 1)
	s := NewSphere()
	for _, test := range tests {
		ray := NewRay(test.origin, direction)
		xs := s.Intersect(ray)
		actual := floatSlice(xs)
		if !sliceEquals(actual, test.expected) {
			t.Errorf("expected intersection t %+v for ray with origin %+v, got %+v", test.expected, test.origin, actual)
		}
		for _, x := range xs {
			if !x.Shape.Equals(s) {
				t.Errorf("expected intersection object %+v for ray with origin %+v, got %+v", Sphere(s), test.origin, x.Shape)
			}
		}
	}
}

func TestRayTransform(t *testing.T) {
	tests := []struct {
		transform Matrix
		expected  Ray
	}{
		{Translate(3, 4, 5), NewRay(NewPoint(4, 6, 8), NewVector(0, 1, 0))},
		{Scale(2, 3, 4), NewRay(NewPoint(2, 6, 12), NewVector(0, 3, 0))},
	}
	ray := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	for _, test := range tests {
		actual := ray.Transform4(test.transform)
		if !test.expected.Equals(actual) {
			t.Errorf("transform using %+v expected %+v, got %+v", test.transform, test.expected, actual)
		}
	}
}
