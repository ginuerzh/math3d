package m32

import (
//"math"
)

type Mat2 [4]float32

func Ident2() Mat2 {
	return Mat2{
		1, 0,
		0, 1,
	}
}

func (m Mat2) At(row, col int) float32 {
	return m[col*2+row]
}

func (m *Mat2) Set(row, col int, v float32) {
	m[col*2+row] = v
}

func (m Mat2) Add(m2 Mat2) Mat2 {
	return Mat2{
		m[0] + m2[0], m[1] + m2[1],
		m[2] + m2[2], m[3] + m2[3],
	}
}

func (m Mat2) Sub(m2 Mat2) Mat2 {
	return Mat2{
		m[0] - m2[0], m[1] - m2[1],
		m[2] - m2[2], m[3] - m2[3],
	}
}

func (m Mat2) Col(col int) Vec2 {
	return Vec2{m[2*col], m[2*col+1]}
}

func (m Mat2) Row(row int) Vec2 {
	return Vec2{m[row], m[row+2]}
}
