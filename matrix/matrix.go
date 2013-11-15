// matrix
package matrix

import (
	"bytes"
	"fmt"
	"github.com/ginuerzh/math3d/vector"
)

type Matrix []vector.Vector

func NewMatrix(row, col int) Matrix {
	m := make([]vector.Vector, col)
	for i, _ := range m {
		m[i] = vector.NewVector(row)
	}
	return m
}

func NewIdentityMatrix(row int) Matrix {
	m := NewMatrix(row, row)
	for i := 1; i <= m.Cols(); i++ {
		m.Column(i).Set(i, 1)
	}
	return m
}

// matrix m multiply by scale, return a new matrix
func MultiSM(scale float64, m Matrix) Matrix {
	return m.Fork().MultiS(scale)
}

func MultiMM(m Matrix, ms ...Matrix) Matrix {
	nm := m
	for _, mm := range ms {
		nm = nm.multiM(mm)
	}

	return nm
}

func (m Matrix) Init(cols ...vector.Vector) {
	for i, _ := range cols {
		copy(m.Column(i+1), cols[i])
	}
}

func (m Matrix) InitColumn(col int, v vector.Vector) {
	if col <= 0 || col > m.Cols() {
		return
	}

	m.Column(col).InitV(v)
}

func (m Matrix) InitRow(row int, v vector.Vector) {
	if row <= 0 || row > m.Rows() {
		return
	}

	for i := 1; i <= m.Cols(); i++ {
		if v.Dim() < i {
			break
		}
		m.Set(row, i, v.Get(i))
	}
}

func (m Matrix) Get(row, col int) float64 {
	if row <= 0 || row > m.Rows() || col <= 0 || col > m.Cols() {
		return 0
	}
	return m.Column(col).Get(row)
}

// set value in [row, col], return old value
func (m Matrix) Set(row, col int, value float64) float64 {
	if row <= 0 || row > m.Rows() || col <= 0 || col > m.Cols() {
		return 0
	}

	return m.Column(col).Set(row, value)
}

func (m Matrix) Cols() int {
	return len(m)
}

func (m Matrix) Rows() int {
	return len(m[0])
}

// return column col in matrix, col starts from 1,
// NOTE: the returned vector is a reference to the row vector in matrix m
func (m Matrix) Column(col int) vector.Vector {
	if col > 0 && col <= len(m) {
		return m[col-1]
	}
	return nil
}

// return row 'row' in matrix, row starts from 1,
// NOTE: the returned vector is a new vector, not a reference to the vector in matrix
func (m Matrix) Row(row int) vector.Vector {
	if row > 0 && row <= m.Rows() {
		v := vector.NewVector(m.Cols())
		for i := 1; i <= m.Cols(); i++ {
			v.Set(i, m.Column(i).Get(row))
		}
		return v
	}

	return nil
}

func (m Matrix) Transpose() Matrix {
	tran := NewMatrix(m.Cols(), m.Rows())

	for i := 1; i <= tran.Cols(); i++ {
		tran.InitColumn(i, m.Row(i))
	}
	return tran
}

func (m Matrix) MultiS(scale float64) Matrix {
	for i := 1; i <= m.Cols(); i++ {
		m.Column(i).Multi(scale)
	}
	return m
}

// return a new matrix
func (m Matrix) multiM(m2 Matrix) Matrix {
	if m.Cols() != m2.Rows() {
		return nil
	}
	mm := NewMatrix(m.Rows(), m2.Cols())
	for i := 1; i <= mm.Cols(); i++ {
		col := mm.Column(i)
		for j := 1; j <= col.Dim(); j++ {
			dot, _ := vector.Dot(m.Row(j), m2.Column(i))
			col.Set(j, dot)
		}
	}

	return mm
}

func (m Matrix) Fork() Matrix {
	mx := NewMatrix(m.Rows(), m.Cols())
	mx.Init(m...)
	return mx
}

func (m Matrix) ToSlice() []float64 {
	s := make([]float64, 0, m.Rows()*m.Cols())

	for i := 1; i <= m.Cols(); i++ {
		s = append(s, m.Column(i)...)
	}

	return s
}

func (m Matrix) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintln(buf)
	for i := 1; i <= m.Rows(); i++ {
		fmt.Fprintln(buf, m.Row(i))
	}

	return buf.String()
}
