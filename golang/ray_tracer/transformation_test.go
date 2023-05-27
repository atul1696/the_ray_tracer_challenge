package ray_tracer

import (
	"math"
	"testing"
)

func TestPointTranslation(t *testing.T) {
	transform := Translate(5, -3, 2)
	p := NewPoint(-3, 4, 5)
	expected := NewPoint(2, 1, 7)
	checkPointTransform(t, &p, &transform, &expected)
}

func TestPointTranslationInverse(t *testing.T) {
	transform, _ := Translate(5, -3, 2).Inverse()
	p := NewPoint(-3, 4, 5)
	expected := NewPoint(-8, 7, 3)
	checkPointTransform(t, &p, &transform, &expected)
}

func TestVectorTranslation(t *testing.T) {
	transform := Translate(5, -3, 2)
	v := NewVector(-3, 4, 5)
	expected := v
	checkVectorTransform(t, &v, &transform, &expected)
}

func TestPointScaling(t *testing.T) {
	transform := Scale(2, 3, 4)
	p := NewPoint(-4, 6, 8)
	expected := NewPoint(-8, 18, 32)
	checkPointTransform(t, &p, &transform, &expected)
}

func TestVectorScaling(t *testing.T) {
	transform := Scale(2, 3, 4)
	v := NewVector(-4, 6, 8)
	expected := NewVector(-8, 18, 32)
	checkVectorTransform(t, &v, &transform, &expected)
}

func TestVectorScalingInverse(t *testing.T) {
	transform, _ := Scale(2, 3, 4).Inverse()
	v := NewVector(-4, 6, 8)
	expected := NewVector(-2, 2, 2)
	checkVectorTransform(t, &v, &transform, &expected)
}

func TestPointReflection(t *testing.T) {
	transform := Scale(-1, 1, 1)
	p := NewPoint(2, 3, 4)
	expected := NewPoint(-2, 3, 4)
	checkPointTransform(t, &p, &transform, &expected)
}
func TestPointRotateX(t *testing.T) {
	tests := []struct {
		r        float64
		expected Point
	}{
		{math.Pi / 4, NewPoint(0, 1/math.Sqrt(2), 1/math.Sqrt(2))},
		{math.Pi / 2, NewPoint(0, 0, 1)},
	}
	p := NewPoint(0, 1, 0)
	for _, test := range tests {
		transform := RotateX(test.r)
		checkPointTransform(t, &p, &transform, &test.expected)
	}
}

func TestPointRotateXInverse(t *testing.T) {
	transform, _ := RotateX(math.Pi / 4).Inverse()
	p := NewPoint(0, 1, 0)
	expected := NewPoint(0, 1/math.Sqrt(2), -1/math.Sqrt(2))
	checkPointTransform(t, &p, &transform, &expected)
}

func TestPointRotateY(t *testing.T) {
	tests := []struct {
		r        float64
		expected Point
	}{
		{math.Pi / 4, NewPoint(1/math.Sqrt(2), 0, 1/math.Sqrt(2))},
		{math.Pi / 2, NewPoint(1, 0, 0)},
	}
	p := NewPoint(0, 0, 1)
	for _, test := range tests {
		transform := RotateY(test.r)
		checkPointTransform(t, &p, &transform, &test.expected)
	}
}

func TestPointRotateZ(t *testing.T) {
	tests := []struct {
		r        float64
		expected Point
	}{
		{math.Pi / 4, NewPoint(-1/math.Sqrt(2), 1/math.Sqrt(2), 0)},
		{math.Pi / 2, NewPoint(-1, 0, 0)},
	}
	p := NewPoint(0, 1, 0)
	for _, test := range tests {
		transform := RotateZ(test.r)
		checkPointTransform(t, &p, &transform, &test.expected)
	}
}

func TestPointShear(t *testing.T) {
	tests := []struct {
		xy       float64
		xz       float64
		yx       float64
		yz       float64
		zx       float64
		zy       float64
		expected Point
	}{
		{1, 0, 0, 0, 0, 0, NewPoint(5, 3, 4)},
		{0, 1, 0, 0, 0, 0, NewPoint(6, 3, 4)},
		{0, 0, 1, 0, 0, 0, NewPoint(2, 5, 4)},
		{0, 0, 0, 1, 0, 0, NewPoint(2, 7, 4)},
		{0, 0, 0, 0, 1, 0, NewPoint(2, 3, 6)},
		{0, 0, 0, 0, 0, 1, NewPoint(2, 3, 7)},
	}
	p := NewPoint(2, 3, 4)
	for _, test := range tests {
		transform := Shear(test.xy, test.xz, test.yx, test.yz, test.zx, test.zy)
		checkPointTransform(t, &p, &transform, &test.expected)
	}
}

func TestTransformChain(t *testing.T) {
	p := NewPoint(1, 0, 1)
	rotx := RotateX(math.Pi / 2)
	s := Scale(5, 5, 5)
	tr := Translate(10, 5, 7)
	expected := NewPoint(15, 0, 7)

	p1, _ := p.Transform(rotx)
	exP1 := NewPoint(1, -1, 0)
	if !p1.Equals(exP1) {
		t.Errorf("%+v rotate x using %+v = %+v, got %+v", p, rotx, exP1, p1)
	}

	p2, _ := p1.Transform(s)
	exP2 := NewPoint(5, -5, 0)
	if !p2.Equals(exP2) {
		t.Errorf("%+v scale using %+v = %+v, got %+v", p1, s, exP2, p2)
	}

	p3, _ := p2.Transform(tr)
	if !p3.Equals(expected) {
		t.Errorf("%+v translate using %+v = %+v, got %+v", p2, tr, exP1, p3)
	}

	transform, _ := tr.Multiply(s)
	transform, _ = transform.Multiply(rotx)
	checkPointTransform(t, &p, &transform, &expected)
}

func TestTransformView(t *testing.T) {
	m := NewMatrix(4, 4)
	m[0][0] = -0.50709
	m[0][1] = 0.50709
	m[0][2] = 0.67612
	m[0][3] = -2.36643
	m[1][0] = 0.76772
	m[1][1] = 0.60609
	m[1][2] = 0.12122
	m[1][3] = -2.82843
	m[2][0] = -0.35857 
	m[2][1] = 0.59761
	m[2][2] = -0.71714
	m[3][3] = float64(1)

	tests := []struct{
		from, to Point
		up Vector
		expected Matrix
	}{
		{NewPoint(0, 0, 0), NewPoint(0, 0, -1), NewVector(0, 1, 0), NewIdentityMatrix(4)},
		{NewPoint(0, 0, 0), NewPoint(0, 0, 1), NewVector(0, 1, 0), Scale(-1, 1, -1)},
		{NewPoint(0, 0, 8), NewPoint(0, 0, 0), NewVector(0, 1, 0), Translate(0, 0, -8)},
		{NewPoint(1, 3, 2), NewPoint(4, -2, 8), NewVector(1, 1, 0), m},
	}
	for _, test := range tests {
		actual := ViewTransform(test.from, test.to, test.up)
		if eq, err := test.expected.Equals(actual); !eq || err != nil {
			if (err != nil) {
				t.Errorf("view transfrom using from: %+v, to: %+v, up: %+v got error: %+v", test.from, test.to, test.up, err)
			} else {
				t.Errorf("view transfrom using from: %+v, to: %+v, up: %+v expected %+v got %+v", test.from, test.to, test.up, test.expected, actual)
			}
		}
	}
}

func checkPointTransform(t *testing.T, p *Point, transform *Matrix, expected *Point) {
	if actual, err := p.Transform(*transform); !(expected.Equals(actual) && err == nil) {
		if err != nil {
			t.Errorf("%+v.Transform() = %+v, got err: %+v", *p, *expected, err)
		} else {
			t.Errorf("%+v.Transform() = %+v, got %+v", *p, *expected, actual)
		}
	}
}

func checkVectorTransform(t *testing.T, v *Vector, transform *Matrix, expected *Vector) {
	if actual, err := v.Transform(*transform); !(expected.Equals(actual) && err == nil) {
		if err != nil {
			t.Errorf("%+v.Transform() = %+v, got err: %+v", *v, *expected, err)
		} else {
			t.Errorf("%+v.Transform() = %+v, got %+v", *v, *expected, actual)
		}
	}
}
