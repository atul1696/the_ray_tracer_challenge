package ray_tracer

import (
	"math"
	"testing"
)

func TestCameraPixelSize(t *testing.T) {
	tests := []struct{
		width, height int
	}{
		{125, 200},
		{200, 125},
	}
	expected := 0.01
	for _, test := range tests {
		c := NewCamera(test.width, test.height, math.Pi/2)
		actual := c.PixelSize()
		if !almostEqual(expected, actual) {
			t.Errorf("camera %d x %d expected %f got %f", test.width, test.height, expected, actual)
		}
	}
}
func TestCameraPixelRay(t *testing.T) {
	r2f := 1 / math.Sqrt(2)
	tests := []struct{
		px, py int
		transform Matrix
		expected Ray
	}{
		{100, 50, NewIdentityMatrix(4), NewRay(NewPoint(0, 0, 0), NewVector(0, 0, -1))},
		{0, 0, NewIdentityMatrix(4), NewRay(NewPoint(0, 0, 0), NewVector(0.66519, 0.33259, -0.66851))},
		{100, 50, RotateY(math.Pi/4).Multiply4(Translate(0, -2, 5)), NewRay(NewPoint(0, 2, -5), NewVector(r2f, 0, -r2f))},
	}
	for _, test := range tests {
		c := NewCamera(201, 101, math.Pi/2)
		c.SetTransform(test.transform)
		actual := c.PixelRay(test.px, test.py)
		if !test.expected.Equals(actual) {
			t.Errorf("camera %+v expected %f got %f", c, test.expected, actual)
		}
	}
}

func TestCameraCanvasRender(t *testing.T) {
	w := NewDefaultWorld()
	c := NewCamera(11, 11, math.Pi / 2)
	c.SetTransform(ViewTransform(NewPoint(0, 0, -5), NewPoint(0, 0, 0), NewVector(0, 1, 0)))
	image := c.Render(w)
	expected := NewColor(0.38066, 0.47583, 0.2855)
	actual := image[5][5]
	if !expected.Equals(actual) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
