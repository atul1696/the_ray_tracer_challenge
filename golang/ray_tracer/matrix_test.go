package ray_tracer

import (
	"fmt"
	"testing"
)

func TestMatrixEquals(t *testing.T) {
	m1 := NewMatrix(4, 4)
	m1[0][0] = float64(1)
	m1[0][1] = float64(2)
	m1[0][2] = float64(3)
	m1[0][3] = float64(4)
	m1[1][0] = float64(5)
	m1[1][1] = float64(6)
	m1[1][2] = float64(7)
	m1[1][3] = float64(8)
	m1[2][0] = float64(9)
	m1[2][1] = float64(8)
	m1[2][2] = float64(7)
	m1[2][3] = float64(6)
	m1[3][0] = float64(5)
	m1[3][1] = float64(4)
	m1[3][2] = float64(3)
	m1[3][3] = float64(2)

	m2 := NewMatrix(4, 4)
	m2[0][0] = float64(1)
	m2[0][1] = float64(2)
	m2[0][2] = float64(3)
	m2[0][3] = float64(4)
	m2[1][0] = float64(5)
	m2[1][1] = float64(6)
	m2[1][2] = float64(7)
	m2[1][3] = float64(8)
	m2[2][0] = float64(9)
	m2[2][1] = float64(8)
	m2[2][2] = float64(7)
	m2[2][3] = float64(6)
	m2[3][0] = float64(5)
	m2[3][1] = float64(4)
	m2[3][2] = float64(3)
	m2[3][3] = float64(2)

	if e, err := m1.Equals(m2); !(e && err == nil) {
		t.Error("matrices are equal, but got not equals")
	}

	m3 := NewMatrix(4, 4)
	m3[0][0] = float64(2)
	m3[0][1] = float64(3)
	m3[0][2] = float64(4)
	m3[0][3] = float64(5)
	m3[1][0] = float64(6)
	m3[1][1] = float64(7)
	m3[1][2] = float64(8)
	m3[1][3] = float64(9)
	m3[2][0] = float64(8)
	m3[2][1] = float64(7)
	m3[2][2] = float64(6)
	m3[2][3] = float64(5)
	m3[3][0] = float64(4)
	m3[3][1] = float64(3)
	m3[3][2] = float64(2)
	m3[3][3] = float64(1)

	if e, err := m1.Equals(m3); !(!e && err == nil) {
		t.Error("matrices are not equal, but got equals")
	}
}

func TestMatrixTranspose(t *testing.T) {
	m1 := NewMatrix(4, 4)
	m1[0][0] = float64(0)
	m1[0][1] = float64(9)
	m1[0][2] = float64(3)
	m1[0][3] = float64(0)
	m1[1][0] = float64(9)
	m1[1][1] = float64(8)
	m1[1][2] = float64(0)
	m1[1][3] = float64(8)
	m1[2][0] = float64(1)
	m1[2][1] = float64(8)
	m1[2][2] = float64(5)
	m1[2][3] = float64(3)
	m1[3][0] = float64(0)
	m1[3][1] = float64(0)
	m1[3][2] = float64(5)
	m1[3][3] = float64(8)

	m2 := NewMatrix(4, 4)
	m2[0][0] = float64(0)
	m2[0][1] = float64(9)
	m2[0][2] = float64(1)
	m2[0][3] = float64(0)
	m2[1][0] = float64(9)
	m2[1][1] = float64(8)
	m2[1][2] = float64(8)
	m2[1][3] = float64(0)
	m2[2][0] = float64(3)
	m2[2][1] = float64(0)
	m2[2][2] = float64(5)
	m2[2][3] = float64(5)
	m2[3][0] = float64(0)
	m2[3][1] = float64(8)
	m2[3][2] = float64(3)
	m2[3][3] = float64(8)

	tr := m1.Transpose()
	if e, err := tr.Equals(m2); !(e && err == nil) {
		t.Error("matrices are equal, but got not equals")
	}
}

func TestMatrixMultiply(t *testing.T) {
	m1 := NewMatrix(4, 4)
	m1[0][0] = float64(1)
	m1[0][1] = float64(2)
	m1[0][2] = float64(3)
	m1[0][3] = float64(4)
	m1[1][0] = float64(5)
	m1[1][1] = float64(6)
	m1[1][2] = float64(7)
	m1[1][3] = float64(8)
	m1[2][0] = float64(9)
	m1[2][1] = float64(8)
	m1[2][2] = float64(7)
	m1[2][3] = float64(6)
	m1[3][0] = float64(5)
	m1[3][1] = float64(4)
	m1[3][2] = float64(3)
	m1[3][3] = float64(2)

	m2 := NewMatrix(4, 4)
	m2[0][0] = float64(-2)
	m2[0][1] = float64(1)
	m2[0][2] = float64(2)
	m2[0][3] = float64(3)
	m2[1][0] = float64(3)
	m2[1][1] = float64(2)
	m2[1][2] = float64(1)
	m2[1][3] = float64(-1)
	m2[2][0] = float64(4)
	m2[2][1] = float64(3)
	m2[2][2] = float64(6)
	m2[2][3] = float64(5)
	m2[3][0] = float64(1)
	m2[3][1] = float64(2)
	m2[3][2] = float64(7)
	m2[3][3] = float64(8)

	p := NewMatrix(4, 4)
	p[0][0] = float64(20)
	p[0][1] = float64(22)
	p[0][2] = float64(50)
	p[0][3] = float64(48)
	p[1][0] = float64(44)
	p[1][1] = float64(54)
	p[1][2] = float64(114)
	p[1][3] = float64(108)
	p[2][0] = float64(40)
	p[2][1] = float64(58)
	p[2][2] = float64(110)
	p[2][3] = float64(102)
	p[3][0] = float64(16)
	p[3][1] = float64(26)
	p[3][2] = float64(46)
	p[3][3] = float64(42)

	prod, err1 := m1.Multiply(m2)
	if err1 != nil {
		t.Error("error when multiplying compatible matrices")
	}
	if e, err2 := p.Equals(prod); !(e && err2 == nil) {
		t.Error("matrices are equal, but got not equals")
	}
}

func TestIdentityMatrix(t *testing.T) {
	dim := 4
	expected := NewMatrix(dim, dim)
	expected[0][0] = float64(1)
	expected[1][1] = float64(1)
	expected[2][2] = float64(1)
	expected[3][3] = float64(1)

	actual := NewIdentityMatrix(dim)
	if e, err := actual.Equals(expected); !(e && err == nil) {
		t.Error("matrices are equal, but got not equals")
	}

	m1 := NewMatrix(4, 4)
	m1[0][0] = float64(1)
	m1[0][1] = float64(2)
	m1[0][2] = float64(3)
	m1[0][3] = float64(4)
	m1[1][0] = float64(5)
	m1[1][1] = float64(6)
	m1[1][2] = float64(7)
	m1[1][3] = float64(8)
	m1[2][0] = float64(9)
	m1[2][1] = float64(8)
	m1[2][2] = float64(7)
	m1[2][3] = float64(6)
	m1[3][0] = float64(5)
	m1[3][1] = float64(4)
	m1[3][2] = float64(3)
	m1[3][3] = float64(2)

	identity := actual
	prod, err1 := identity.Multiply(m1)
	if err1 != nil {
		t.Error("error when multiplying compatible matrices")
	}
	if e, err2 := m1.Equals(prod); !(e && err2 == nil) {
		t.Error("matrices are equal, but got not equals")
	}
}

func TestMatrixDeterminant1(t *testing.T) {
	m := NewMatrix(2, 2)
	m[0][0] = 1
	m[0][1] = 5
	m[1][0] = -3
	m[1][1] = 2

	expected := float64(17)
	actual := m.Determinant()
	if !almostEqual(actual, expected) {
		t.Errorf("m.Determinant() = %f, got %f", expected, actual)
	}
}

func TestSubMatrix1(t *testing.T) {
	m1 := NewMatrix(3, 3)
	m1[0][0] = float64(1)
	m1[0][1] = float64(5)
	m1[0][2] = float64(0)
	m1[1][0] = float64(-3)
	m1[1][1] = float64(2)
	m1[1][2] = float64(7)
	m1[2][0] = float64(0)
	m1[2][1] = float64(6)
	m1[2][2] = float64(-3)

	expected := NewMatrix(2, 2)
	expected[0][0] = float64(-3)
	expected[0][1] = float64(2)
	expected[1][0] = float64(0)
	expected[1][1] = float64(6)

	actual := m1.SubMatrix(0, 2)
	if eq, err := expected.Equals(actual); !(eq && err == nil) {
		t.Error(fmt.Errorf("submatrix not equal %v != %v", expected, actual))
	}
}

func TestSubMatrix2(t *testing.T) {
	m1 := NewMatrix(4, 4)
	m1[0][0] = float64(-6)
	m1[0][1] = float64(1)
	m1[0][2] = float64(1)
	m1[0][3] = float64(6)
	m1[1][0] = float64(-8)
	m1[1][1] = float64(5)
	m1[1][2] = float64(8)
	m1[1][3] = float64(6)
	m1[2][0] = float64(-1)
	m1[2][1] = float64(0)
	m1[2][2] = float64(8)
	m1[2][3] = float64(2)
	m1[3][0] = float64(-7)
	m1[3][1] = float64(1)
	m1[3][2] = float64(-1)
	m1[3][3] = float64(1)

	expected := NewMatrix(3, 3)
	expected[0][0] = float64(-6)
	expected[0][1] = float64(1)
	expected[0][2] = float64(6)
	expected[1][0] = float64(-8)
	expected[1][1] = float64(8)
	expected[1][2] = float64(6)
	expected[2][0] = float64(-7)
	expected[2][1] = float64(-1)
	expected[2][2] = float64(1)

	actual := m1.SubMatrix(2, 1)
	if eq, err := expected.Equals(actual); !(eq && err == nil) {
		t.Error(fmt.Errorf("submatrix not equal %v != %v", expected, actual))
	}
}

func TestMatrixMinor(t *testing.T) {
	m := NewMatrix(3, 3)
	m[0][0] = 3
	m[0][1] = 5
	m[0][2] = 0
	m[1][0] = 2
	m[1][1] = -1
	m[1][2] = -7
	m[2][0] = 6
	m[2][1] = -1
	m[2][2] = 5

	expected := float64(25)
	actual := m.Minor(1, 0)
	if !almostEqual(expected, actual) {
		t.Errorf("m.Minor() = %f, got %f", expected, actual)
	}
}

func TestMatrixCofactor(t *testing.T) {
	m := NewMatrix(3, 3)
	m[0][0] = 3
	m[0][1] = 5
	m[0][2] = 0
	m[1][0] = 2
	m[1][1] = -1
	m[1][2] = -7
	m[2][0] = 6
	m[2][1] = -1
	m[2][2] = 5

	var tests = []struct {
		row               int
		column            int
		expected_minor    float64
		expected_cofactor float64
	}{
		{0, 0, -12, -12},
		{1, 0, 25, -25},
	}
	for _, test := range tests {
		if actual_minor := m.Minor(test.row, test.column); !almostEqual(test.expected_minor, actual_minor) {
			t.Errorf("m.Minor(%d, %d) = %f, got %f", test.row, test.column, test.expected_minor, actual_minor)
		}
		if actual_cofactor := m.Cofactor(test.row, test.column); !almostEqual(test.expected_cofactor, actual_cofactor) {
			t.Errorf("m.Cofactor(%d, %d) = %f, got %f", test.row, test.column, test.expected_cofactor, actual_cofactor)
		}
	}
}

func TestMatrixDeterminant2(t *testing.T) {
	m := NewMatrix(3, 3)
	m[0][0] = float64(1)
	m[0][1] = float64(2)
	m[0][2] = float64(6)
	m[1][0] = float64(-5)
	m[1][1] = float64(8)
	m[1][2] = float64(-4)
	m[2][0] = float64(2)
	m[2][1] = float64(6)
	m[2][2] = float64(4)

	var cofactor_tests = []struct {
		row      int
		column   int
		expected float64
	}{
		{0, 0, 56},
		{0, 1, 12},
		{0, 2, -46},
	}
	for _, test := range cofactor_tests {
		if actual := m.Cofactor(test.row, test.column); !almostEqual(test.expected, actual) {
			t.Errorf("m.Cofactor(%d, %d) = %f, got %f", test.row, test.column, test.expected, actual)
		}
	}

	expected := float64(-196)
	actual := m.Determinant()
	if !almostEqual(actual, expected) {
		t.Errorf("m.Determinant() = %f, got %f", expected, actual)
	}
}

func TestMatrixDeterminant3(t *testing.T) {
	m := NewMatrix(4, 4)
	m[0][0] = float64(-2)
	m[0][1] = float64(-8)
	m[0][2] = float64(3)
	m[0][3] = float64(5)
	m[1][0] = float64(-3)
	m[1][1] = float64(1)
	m[1][2] = float64(7)
	m[1][3] = float64(3)
	m[2][0] = float64(1)
	m[2][1] = float64(2)
	m[2][2] = float64(-9)
	m[2][3] = float64(6)
	m[3][0] = float64(-6)
	m[3][1] = float64(7)
	m[3][2] = float64(7)
	m[3][3] = float64(-9)

	var cofactor_tests = []struct {
		row      int
		column   int
		expected float64
	}{
		{0, 0, 690},
		{0, 1, 447},
		{0, 2, 210},
		{0, 3, 51},
	}
	for _, test := range cofactor_tests {
		if actual := m.Cofactor(test.row, test.column); !almostEqual(test.expected, actual) {
			t.Errorf("m.Cofactor(%d, %d) = %f, got %f", test.row, test.column, test.expected, actual)
		}
	}

	expected := float64(-4071)
	actual := m.Determinant()
	if !almostEqual(actual, expected) {
		t.Errorf("m.Determinant() = %f, got %f", expected, actual)
	}
}

func TestMatrixInverse1(t *testing.T) {
	m := NewMatrix(4, 4)
	m[0][0] = float64(8)
	m[0][1] = float64(-5)
	m[0][2] = float64(9)
	m[0][3] = float64(2)
	m[1][0] = float64(7)
	m[1][1] = float64(5)
	m[1][2] = float64(6)
	m[1][3] = float64(1)
	m[2][0] = float64(-6)
	m[2][1] = float64(0)
	m[2][2] = float64(9)
	m[2][3] = float64(6)
	m[3][0] = float64(-3)
	m[3][1] = float64(0)
	m[3][2] = float64(-9)
	m[3][3] = float64(-4)

	expected := NewMatrix(4, 4)
	expected[0][0] = float64(-0.15385)
	expected[0][1] = float64(-0.15385)
	expected[0][2] = float64(-0.28205)
	expected[0][3] = float64(-0.53846)
	expected[1][0] = float64(-0.07692)
	expected[1][1] = float64(0.12308)
	expected[1][2] = float64(0.02564)
	expected[1][3] = float64(0.03077)
	expected[2][0] = float64(0.35897)
	expected[2][1] = float64(0.35897)
	expected[2][2] = float64(0.43590)
	expected[2][3] = float64(0.92308)
	expected[3][0] = float64(-0.69231)
	expected[3][1] = float64(-0.69231)
	expected[3][2] = float64(-0.76923)
	expected[3][3] = float64(-1.92308)

	actual, err1 := m.Inverse()
	if err1 != nil {
		t.Error("can invert matrix")
	}
	if eq, err2 := expected.Equals(actual); !(eq && err2 == nil) {
		t.Errorf("m.inverse() = %v, got %v", expected, actual)
	}
}

func TestMatrixInverse2(t *testing.T) {
	m := NewMatrix(4, 4)
	m[0][0] = float64(9)
	m[0][1] = float64(3)
	m[0][2] = float64(0)
	m[0][3] = float64(9)
	m[1][0] = float64(-5)
	m[1][1] = float64(-2)
	m[1][2] = float64(-6)
	m[1][3] = float64(-3)
	m[2][0] = float64(-4)
	m[2][1] = float64(9)
	m[2][2] = float64(6)
	m[2][3] = float64(4)
	m[3][0] = float64(-7)
	m[3][1] = float64(6)
	m[3][2] = float64(6)
	m[3][3] = float64(2)

	expected := NewMatrix(4, 4)
	expected[0][0] = float64(-0.04074)
	expected[0][1] = float64(-0.07778)
	expected[0][2] = float64(0.14444)
	expected[0][3] = float64(-0.22222)
	expected[1][0] = float64(-0.07778)
	expected[1][1] = float64(0.03333)
	expected[1][2] = float64(0.36667)
	expected[1][3] = float64(-0.33333)
	expected[2][0] = float64(-0.02901)
	expected[2][1] = float64(-0.14630)
	expected[2][2] = float64(-0.10926)
	expected[2][3] = float64(0.12963)
	expected[3][0] = float64(0.17778)
	expected[3][1] = float64(0.06667)
	expected[3][2] = float64(-0.26667)
	expected[3][3] = float64(0.33333)

	actual, err1 := m.Inverse()
	if err1 != nil {
		t.Error("can invert matrix")
	}
	if eq, err2 := expected.Equals(actual); !(eq && err2 == nil) {
		t.Errorf("m.inverse() = %v, got %v", expected, actual)
	}
}

func TestMatrixInverseMultiply(t *testing.T) {
	m1 := NewMatrix(4, 4)
	m1[0][0] = float64(3)
	m1[0][1] = float64(-9)
	m1[0][2] = float64(7)
	m1[0][3] = float64(3)
	m1[1][0] = float64(3)
	m1[1][1] = float64(-8)
	m1[1][2] = float64(2)
	m1[1][3] = float64(-9)
	m1[2][0] = float64(-4)
	m1[2][1] = float64(4)
	m1[2][2] = float64(4)
	m1[2][3] = float64(1)
	m1[3][0] = float64(-6)
	m1[3][1] = float64(5)
	m1[3][2] = float64(-1)
	m1[3][3] = float64(1)

	m2 := NewMatrix(4, 4)
	m2[0][0] = float64(8)
	m2[0][1] = float64(2)
	m2[0][2] = float64(2)
	m2[0][3] = float64(2)
	m2[1][0] = float64(3)
	m2[1][1] = float64(-1)
	m2[1][2] = float64(7)
	m2[1][3] = float64(0)
	m2[2][0] = float64(7)
	m2[2][1] = float64(0)
	m2[2][2] = float64(5)
	m2[2][3] = float64(4)
	m2[3][0] = float64(6)
	m2[3][1] = float64(-2)
	m2[3][2] = float64(0)
	m2[3][3] = float64(5)

	p1, err := m1.Multiply(m2)
	if err != nil {
		t.Error("error when multiplying compatible matrices")
	}
	inv, err1 := m2.Inverse()
	if err1 != nil {
		t.Error("can invert matrix")
	}
	actual, err2 := p1.Multiply(inv)
	if err2 != nil {
		t.Error("error when multiplying compatible matrices")
	}
	if eq, err3 := m1.Equals(actual); !(eq && err3 == nil) {
		t.Errorf("inverse product not equal, %v != %v", m1, actual)
	}
}
