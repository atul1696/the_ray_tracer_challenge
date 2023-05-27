package ray_tracer

import "testing"

func TestComputationOutside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	expectedT := 4.0
	x := NewIntersection(expectedT, s)
	inside := false

	actual := PrepareComputation(x, r)
	if actual.T != expectedT {
		t.Errorf("t: expected %f, got %f", expectedT, actual.T)
	}
	if !s.Equals(actual.Object) {
		t.Errorf("object: expected %+v, got %+v", s, actual.Object)
	}
	expectedPoint := NewPoint(0, 0, -1)
	if !expectedPoint.Equals(actual.Point) {
		t.Errorf("point: expected %+v, got %+v", expectedPoint, actual.Point)
	}
	expectedEyeV := NewVector(0, 0, -1)
	if !expectedEyeV.Equals(actual.EyeV) {
		t.Errorf("eyeV: expected %+v, got %+v", expectedEyeV, actual.EyeV)
	}
	expectedNormalV := NewVector(0, 0, -1)
	if !expectedNormalV.Equals(actual.NormalV) {
		t.Errorf("normalV: expected %+v, got %+v", expectedNormalV, actual.NormalV)
	}
	if actual.Inside != inside {
		t.Error("inside: expected false, got true")
	}
}

func TestComputationInside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s := NewSphere()
	expectedT := 1.0
	x := NewIntersection(expectedT, s)
	inside := true

	actual := PrepareComputation(x, r)
	if actual.T != expectedT {
		t.Errorf("t: expected %f, got %f", expectedT, actual.T)
	}
	if !s.Equals(actual.Object) {
		t.Errorf("object: expected %+v, got %+v", s, actual.Object)
	}
	expectedPoint := NewPoint(0, 0, 1)
	if !expectedPoint.Equals(actual.Point) {
		t.Errorf("point: expected %+v, got %+v", expectedPoint, actual.Point)
	}
	expectedEyeV := NewVector(0, 0, -1)
	if !expectedEyeV.Equals(actual.EyeV) {
		t.Errorf("eyeV: expected %+v, got %+v", expectedEyeV, actual.EyeV)
	}
	expectedNormalV := NewVector(0, 0, -1)
	if !expectedNormalV.Equals(actual.NormalV) {
		t.Errorf("normalV: expected %+v, got %+v", expectedNormalV, actual.NormalV)
	}
	if actual.Inside != inside {
		t.Error("inside: expected true, got false")
	}
}
