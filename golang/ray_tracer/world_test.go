package ray_tracer

import (
	"testing"
)

func TestWorldIntersect(t *testing.T) {
	world := NewDefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	expected := []float64{4, 4.5, 5.5, 6}
	xs := world.Intersect(ray)
	actual := floatSlice(xs)
	if !sliceEquals(expected, actual) {
		t.Errorf("default world intersect by ray %+v, expected time slice %+v got %+v", ray, expected, actual)
	}
}

func TestWorldShadingOutside(t *testing.T) {
	world := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	x := NewIntersection(4, world.Shapes[0])
	c := PrepareComputation(x, r)
	actual := world.ShadeHit(c)
	expected := NewColor(0.38066, 0.47583, 0.2855)
	if !expected.Equals(actual) {
		t.Errorf("%+v", world.Shapes[0])
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestWorldShadingInside(t *testing.T) {
	world := NewDefaultWorld()
	world.SetLight(NewPointLight(NewPoint(0, 0.25, 0), NewColor(1, 1, 1)))
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	x := NewIntersection(0.5, world.Shapes[1])
	c := PrepareComputation(x, r)
	actual := world.ShadeHit(c)
	expected := NewColor(0.90498, 0.90498, 0.90498)
	if !expected.Equals(actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestWorldColorAtOutside(t *testing.T) {
	p := NewPoint(0, 0, -5)
	w := NewDefaultWorld()
	tests := []struct {
		dir      Vector
		expected Color
	}{
		{NewVector(0, 1, 0), NewColor(0, 0, 0)},
		{NewVector(0, 0, 1), NewColor(0.38066, 0.47583, 0.2855)},
	}
	for _, test := range tests {
		r := NewRay(p, test.dir)
		actual := w.ColorAt(r)
		if !test.expected.Equals(actual) {
			t.Errorf("for ray: %+v, expected %+v got %+v", r, test.expected, actual)
		}
	}
}

func TestWorldColorAtInside(t *testing.T) {
	w := NewDefaultWorld()
	for i := range w.Shapes {
		w.Shapes[i].Material.SetAmbient(1.0)
	}
	r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))
	expected := w.Shapes[1].Material.Color
	actual := w.ColorAt(r)
	if !expected.Equals(actual) {
		t.Errorf("for ray: %+v, expected %+v got %+v", r, expected, actual)
	}
}

func TestWorldPointInShadow(t *testing.T) {
	w := NewDefaultWorld()
	tests := []struct {
		point    Point
		expected bool
	}{
		{NewPoint(0, 10, 0), false},
		{NewPoint(10, -10, 10), true},
		{NewPoint(-20, 20, -20), false},
		{NewPoint(-2, 2, -2), false},
	}

	for _, test := range tests {
		actual := w.isInShadow(test.point)
		if test.expected != actual {
			t.Errorf("for point: %+v, expected %+v got %+v", test.point, test.expected, actual)
		}
	}
}

func TestWorldShadowShade(t *testing.T) {
	w := NewWorld()
	w.SetLight(NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1)))
	s1 := NewSphere()
	w.AddShape(s1)
	s2 := NewSphere()
	s2.SetTransform(Translate(0, 0, 10))
	w.AddShape(s2)
	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	x := NewIntersection(4, s2)
	comp := PrepareComputation(x, r)
	expected := NewColor(0.1, 0.1, 0.1)
	actual := w.ShadeHit(comp)

	if !expected.Equals(actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
