package ray_tracer

import (
	"errors"
	"fmt"
)

type Matrix [][]float64

func NewTuple(r int) Matrix {
	return NewMatrix(r, 1)
}

func NewMatrix(r, c int) Matrix {
	mat := make([][]float64, r)
	for i := 0; i < r; i++ {
		mat[i] = make([]float64, c)
	}
	return mat
}

func NewIdentityMatrix(dim int) Matrix {
	mat := make([][]float64, dim)
	for i := 0; i < dim; i++ {
		mat[i] = make([]float64, dim)
		mat[i][i] = float64(1)
	}
	return mat
}

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Columns() int {
	return len(m[0])
}

func (m Matrix) Equals(m1 Matrix) (bool, error) {
	if m.Rows() != m1.Rows() || m.Columns() != m1.Columns() {
		return false, fmt.Errorf("cannot compare matrices of different dimensions: [%d, %d] != [%d, %d]", m.Rows(), m.Columns(), m1.Rows(), m1.Columns())
	}
	equals := true
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Columns(); j++ {
			equals = equals && almostEqual(m[i][j], m1[i][j])
		}
	}
	return equals, nil
}

func (m Matrix) Transpose() Matrix {
	rows, cols := m.Rows(), m.Columns()
	m1 := NewMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m1[i][j] = m[j][i]
		}
	}
	return m1
}

func (m Matrix) Multiply(m1 Matrix) (Matrix, error) {
	if m.Rows() != m1.Columns() {
		return nil, fmt.Errorf("cannot multiply matrices: number of rows %d != number of columns %d]", m1.Rows(), m.Columns())
	}
	r, c := m.Rows(), m1.Columns()
	l := m.Columns()
	prod := NewMatrix(m.Rows(), m1.Columns())
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for k := 0; k < l; k++ {
				prod[i][j] += m[i][k] * m1[k][j]
			}
		}
	}
	return prod, nil
}

func (m Matrix) Determinant() float64 {
	det := float64(0)
	if m.Rows() == 2 {
		det = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	} else {
		for j := 0; j < m.Columns(); j++ {
			det += m[0][j] * m.Cofactor(0, j)
		}
	}
	return det
}

func (m Matrix) SubMatrix(row, column int) Matrix {
	r, c := m.Rows(), m.Columns()
	sub := NewMatrix(r-1, c-1)

	u, v := 0, 0
	for i := 0; i < r; i++ {
		if i == row {
			continue
		}
		for j := 0; j < c; j++ {
			if j != column {
				sub[u][v] = m[i][j]
				v += 1
			}
		}
		u += 1
		v = 0
	}
	return sub
}

func (m Matrix) Minor(row, column int) float64 {
	return m.SubMatrix(row, column).Determinant()
}

func (m Matrix) Cofactor(row, column int) float64 {
	cofactor := m.Minor(row, column)
	if (row+column)&1 == 1 {
		cofactor *= -1.0
	}
	return cofactor
}

func (m Matrix) Inverse() (Matrix, error) {
	r, c := m.Rows(), m.Columns()
	if r != c {
		return nil, errors.New("cannot invert non-square matrix")
	}
	det := m.Determinant()
	if det == 0 {
		return nil, errors.New("matrix is not invertible")
	}
	inverse := NewMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			inverse[j][i] = m.Cofactor(i, j) / det
		}
	}
	return inverse, nil
}
