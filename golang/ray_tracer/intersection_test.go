package ray_tracer

import (
	"testing"
)

func TestIntersectionsHit(t *testing.T) {
	s := NewSphere()
	tests := []struct {
		xts      []float64
		ok       bool
		expected Intersection
	}{
		{[]float64{1, 2}, true, NewIntersection(1, s)},
		{[]float64{-1, 2}, true, NewIntersection(2, s)},
		{[]float64{-2, -1}, false, Intersection{}},
		{[]float64{5, 7, -3, 2}, true, NewIntersection(2, s)},
	}
	for _, test := range tests {
		xs := NewIntersections()
		for _, xt := range test.xts {
			xs.append(NewIntersection(xt, s))
		}
		actual, ok := xs.Hit()
		if ok != test.ok && !test.expected.Equals(actual) {
			t.Errorf("test: %+v, expected %+v got %+v", test.xts, test.expected, actual)
		}
	}
}

func TestIntersectionsRaySphere(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	tests := []struct {
		transform Matrix
		expected  []float64
	}{
		{Scale(2, 2, 2), []float64{3, 7}},
		{Translate(5, 0, 0), []float64{}},
	}
	for _, test := range tests {
		s.SetTransform(test.transform)
		xs := s.Intersect(r)
		actual := floatSlice(xs)
		if !sliceEquals(test.expected, actual) {
			t.Errorf("ray %+v intersect with sphere with transform %+v expected t: %+v, got %+v", r, test.transform, test.expected, actual)
		}
	}
}
