// matrix_test
package matrix

import (
	"github.com/ginuerzh/math3d/vector"
	"testing"
)

var (
	m1 = NewMatrix(3, 3)
	m2 = NewMatrix(3, 3)
)

func init() {
	m1.Init(vector.NewVector(3, 1, 2, 3), vector.NewVector(3, 4, 5, 6), vector.NewVector(3, 7, 8, 9))
	m2.Init(vector.NewVector(3, 7, 8, 9), vector.NewVector(3, 4, 5, 6), vector.NewVector(3, 1, 2, 3))
}

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(3, 3)
	m.Init(vector.NewVector(3, 1, 2, 3), vector.NewVector(3, 4, 5, 6), vector.NewVector(3, 7, 8, 9))
	t.Log(m)
}

func TestMultiSM(t *testing.T) {
	m := MultiSM(2, m1)
	t.Log(m1, m)
}

func TestMultiMM(t *testing.T) {
	m := MultiMM(m1, m2)
	t.Log(m1, m2, m)
}

func TestTranspose(t *testing.T) {
	tran := m1.Transpose()
	t.Log(m1, tran)
}

func TestToSlice(t *testing.T) {
	s := m1.ToSlice()
	t.Log(m1, s)
}
