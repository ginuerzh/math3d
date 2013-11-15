// math3d
package math3d

import (
	"github.com/ginuerzh/math3d/matrix"
	"github.com/ginuerzh/math3d/vector"
	"math"
)

// a 4X4 matrix type, just a plain matrix
type Matrix4 struct {
	matrix.Matrix
}

// a 3 components vector type, just a plain vector
type Vector3 struct {
	vector.Vector
}

// a 4 components vector type
type Vector4 struct {
	vector.Vector
}

func NewMatrix4() *Matrix4 {
	return &Matrix4{matrix.NewMatrix(4, 4)}
}

func NewIdentityMatrix4() *Matrix4 {
	return &Matrix4{matrix.NewIdentityMatrix(4)}
}

func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{vector.NewVector(3, x, y, z)}
}

func NewVector4(x, y, z, w float64) *Vector4 {
	return &Vector4{vector.NewVector(4, x, y, z, w)}
}

// Matrix multiply by column vector, result is a new column vector
func (m *Matrix4) MultiV(v *Vector4) *Vector4 {
	m2 := matrix.NewMatrix(4, 1) // column vector
	m2.Init(v.Vector)

	return &Vector4{matrix.MultiMM(m.Matrix, m2).Column(1)}
}

// Row vector multiply by matrix, result is a new row vector
func (v *Vector4) MultiM(m *Matrix4) *Vector4 {
	m2 := matrix.NewMatrix(1, 4)
	m2.InitRow(1, v.Vector)
	return &Vector4{matrix.MultiMM(m2, m.Matrix).Row(1)}
}

func Translate(x, y, z float64) *Matrix4 {
	m := NewIdentityMatrix4()
	col := m.Column(m.Cols())
	col.Set(1, x)
	col.Set(2, y)
	col.Set(3, z)

	return m
}

func TranslateV(v *Vector3) *Matrix4 {
	return Translate(v.Get(1), v.Get(2), v.Get(3))
}

func Rotate(angle float64, x, y, z float64) *Matrix4 {
	x2 := x * x
	y2 := y * y
	z2 := z * z
	rads := angle * (math.Pi / 180.0)
	c := math.Cos(rads)
	s := math.Sin(rads)
	omc := 1.0 - c

	v1 := NewVector4(x2*omc+c, y*x*omc+z*s, x*z*omc-y*s, 0)
	v2 := NewVector4(x*y*omc-z*s, y2*omc+c, y*z*omc+x*s, 0)
	v3 := NewVector4(x*z*omc+y*s, y*z*omc-x*s, z2*omc+c, 0)
	v4 := NewVector4(0, 0, 0, 1)

	m := NewMatrix4()
	m.Init(v1.Vector, v2.Vector, v3.Vector, v4.Vector)

	return m
}

func RotateV(angle float64, v *Vector3) *Matrix4 {
	return Rotate(angle, v.Get(1), v.Get(2), v.Get(3))
}

func RotateAngle(angleX, angleY, angleZ float64) *Matrix4 {
	m1 := Rotate(angleZ, 0, 0, 1).Matrix
	m2 := Rotate(angleY, 0, 1, 0).Matrix
	m3 := Rotate(angleX, 1, 0, 0).Matrix

	return &Matrix4{matrix.MultiMM(m1, m2, m3)}
}

func Scale(x, y, z float64) *Matrix4 {
	m := NewMatrix4()
	v1 := NewVector4(x, 0, 0, 0)
	v2 := NewVector4(0, y, 0, 0)
	v3 := NewVector4(0, 0, z, 0)
	v4 := NewVector4(0, 0, 0, 1)

	m.Init(v1.Vector, v2.Vector, v3.Vector, v4.Vector)
	return m
}

func ScaleV(v *Vector3) *Matrix4 {
	return Scale(v.Get(1), v.Get(2), v.Get(3))
}

func ScaleX(x float64) *Matrix4 {
	return Scale(x, x, x)
}
