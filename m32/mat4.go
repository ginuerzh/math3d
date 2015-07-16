package m32

import (
	"math"
)

type Mat4 [16]float32

func ZeroMat4() Mat4 {
	return Mat4{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
}

func Ident4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (m Mat4) At(row, col int) float32 {
	return m[col*4+row]
}

func (m *Mat4) Set(row, col int, v float32) {
	m[col*4+row] = v
}

func (m Mat4) Add(m4 Mat4) Mat4 {
	return Mat4{
		m[0] + m4[0], m[1] + m4[1], m[2] + m4[2], m[3] + m4[3],
		m[4] + m4[4], m[5] + m4[5], m[6] + m4[6], m[7] + m4[7],
		m[8] + m4[8], m[9] + m4[9], m[10] + m4[10], m[11] + m4[11],
		m[12] + m4[12], m[13] + m4[13], m[14] + m4[14], m[15] + m4[15],
	}
}

func (m Mat4) Sub(m4 Mat4) Mat4 {
	return Mat4{
		m[0] - m4[0], m[1] - m4[1], m[2] - m4[2], m[3] - m4[3],
		m[4] - m4[4], m[5] - m4[5], m[6] - m4[6], m[7] - m4[7],
		m[8] - m4[8], m[9] - m4[9], m[10] - m4[10], m[11] - m4[11],
		m[12] - m4[12], m[13] - m4[13], m[14] - m4[14], m[15] - m4[15],
	}
}

func (m Mat4) Col(col int) Vec4 {
	return Vec4{m[4*col], m[4*col+1], m[4*col+2], m[4*col+3]}
}

func (m Mat4) Row(row int) Vec4 {
	return Vec4{m[row], m[4+row], m[8+row], m[12+row]}
}

func (m Mat4) Transpose() Mat4 {
	return Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

func (m Mat4) Mul(f float32) Mat4 {
	return Mat4{
		m[0] * f, m[1] * f, m[2] * f, m[3] * f,
		m[4] * f, m[5] * f, m[6] * f, m[7] * f,
		m[8] * f, m[9] * f, m[10] * f, m[11] * f,
		m[12] * f, m[13] * f, m[14] * f, m[15] * f,
	}
}

func (m Mat4) Mul4x1(v Vec4) Vec4 {
	return Vec4{
		m[0]*v[0] + m[4]*v[1] + m[8]*v[2] + m[12]*v[3],
		m[1]*v[0] + m[5]*v[1] + m[9]*v[2] + m[13]*v[3],
		m[2]*v[0] + m[6]*v[1] + m[10]*v[2] + m[14]*v[3],
		m[3]*v[0] + m[7]*v[1] + m[11]*v[2] + m[15]*v[3],
	}
}

func (m Mat4) Mul4(m4 Mat4) Mat4 {
	return Mat4{
		// col 0
		m[0]*m4[0] + m[4]*m4[1] + m[8]*m4[2] + m[12]*m4[3],
		m[1]*m4[0] + m[5]*m4[1] + m[9]*m4[2] + m[13]*m4[3],
		m[2]*m4[0] + m[6]*m4[1] + m[10]*m4[2] + m[14]*m4[3],
		m[3]*m4[0] + m[7]*m4[1] + m[11]*m4[2] + m[15]*m4[3],
		// col 1
		m[0]*m4[4] + m[4]*m4[5] + m[8]*m4[6] + m[12]*m4[7],
		m[1]*m4[4] + m[5]*m4[5] + m[9]*m4[6] + m[13]*m4[7],
		m[2]*m4[4] + m[6]*m4[5] + m[10]*m4[6] + m[14]*m4[7],
		m[3]*m4[4] + m[7]*m4[5] + m[11]*m4[6] + m[15]*m4[7],
		// col 2
		m[0]*m4[8] + m[4]*m4[9] + m[8]*m4[10] + m[12]*m4[11],
		m[1]*m4[8] + m[5]*m4[9] + m[9]*m4[10] + m[13]*m4[11],
		m[2]*m4[8] + m[6]*m4[9] + m[10]*m4[10] + m[14]*m4[11],
		m[3]*m4[8] + m[7]*m4[9] + m[11]*m4[10] + m[15]*m4[11],
		// col 3
		m[0]*m4[12] + m[4]*m4[13] + m[8]*m4[14] + m[12]*m4[15],
		m[1]*m4[12] + m[5]*m4[13] + m[9]*m4[14] + m[13]*m4[15],
		m[2]*m4[12] + m[6]*m4[13] + m[10]*m4[14] + m[14]*m4[15],
		m[3]*m4[12] + m[7]*m4[13] + m[11]*m4[14] + m[15]*m4[15],
	}
}

func (m Mat4) Translate(v Vec3) Mat4 {
	mt := Ident4()
	mt[12] = v[0]
	mt[13] = v[1]
	mt[14] = v[2]
	return mt.Mul4(m)
}

func (m Mat4) Scale(v Vec3) Mat4 {
	ms := Ident4()
	ms[0] = v[0]
	ms[5] = v[1]
	ms[10] = v[2]
	return ms.Mul4(m)
}

func (m Mat4) RotateX(degree float32) Mat4 {
	rad := float64(degree) * RadPerDeg
	mr := Ident4()
	mr[5] = float32(math.Cos(rad))
	mr[9] = float32(-math.Sin(rad))
	mr[6] = float32(math.Sin(rad))
	mr[10] = float32(math.Cos(rad))
	return mr.Mul4(m)
}

func (m Mat4) RotateY(degree float32) Mat4 {
	rad := float64(degree) * RadPerDeg
	mr := Ident4()
	mr[0] = float32(math.Cos(rad))
	mr[8] = float32(math.Sin(rad))
	mr[2] = float32(-math.Sin(rad))
	mr[10] = float32(math.Cos(rad))
	return mr.Mul4(m)
}

func (m Mat4) RotateZ(degree float32) Mat4 {
	rad := float64(degree) * RadPerDeg
	mr := Ident4()
	mr[0] = float32(math.Cos(rad))
	mr[4] = float32(-math.Sin(rad))
	mr[1] = float32(math.Sin(rad))
	mr[5] = float32(math.Cos(rad))
	return mr.Mul4(m)
}
